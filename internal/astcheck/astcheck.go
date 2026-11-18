// Package astcheck implements a small allowlist of syntax constraints for
// exercises whose learning objective is a specific construct (for example
// "use defer" or "declare a constant"). Rules inspect the parsed AST, never
// raw source text, so formatting and comments do not matter.
//
// Rules are written in exercise.json as "name" or "name:argument":
//
//	uses_defer:Func          Func contains a defer statement
//	declares_const:Name      a package-level constant Name exists
//	no_call:pkg.Name         the code never calls pkg.Name (e.g. sort.Ints)
//	calls:pkg.Name           the code calls pkg.Name
//	uses_range:Func          Func contains a for ... range loop
//	uses_switch:Func         Func contains a switch statement
//	chan_params_directional:Func  every channel parameter of Func is send- or receive-only
package astcheck

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

// Rule is a parsed constraint.
type Rule struct {
	Name string
	Arg  string
}

func (r Rule) String() string {
	if r.Arg == "" {
		return r.Name
	}
	return r.Name + ":" + r.Arg
}

type ruleSpec struct {
	needsArg bool
	check    func(p *program, arg string) error
}

var rules = map[string]ruleSpec{
	"uses_defer":              {true, funcContains("a defer statement", func(n ast.Node) bool { _, ok := n.(*ast.DeferStmt); return ok })},
	"uses_range":              {true, funcContains("a for ... range loop", func(n ast.Node) bool { _, ok := n.(*ast.RangeStmt); return ok })},
	"uses_switch":             {true, funcContains("a switch statement", isSwitch)},
	"declares_const":          {true, declaresConst},
	"calls":                   {true, calls(true)},
	"no_call":                 {true, calls(false)},
	"chan_params_directional": {true, chanParamsDirectional},
}

// Parse validates a rule string. It is used by the catalog so unknown rules
// fail curriculum validation instead of grading.
func Parse(s string) (Rule, error) {
	name, arg, _ := strings.Cut(s, ":")
	spec, ok := rules[name]
	if !ok {
		names := make([]string, 0, len(rules))
		for n := range rules {
			names = append(names, n)
		}
		sort.Strings(names)
		return Rule{}, fmt.Errorf("unknown ast rule %q (known: %s)", name, strings.Join(names, ", "))
	}
	if spec.needsArg && arg == "" {
		return Rule{}, fmt.Errorf("ast rule %q needs an argument (%s:...)", name, name)
	}
	if (name == "calls" || name == "no_call") && !strings.Contains(arg, ".") {
		return Rule{}, fmt.Errorf("ast rule %q wants pkg.Name, got %q", name, arg)
	}
	return Rule{Name: name, Arg: arg}, nil
}

// Violation describes a rule the code does not satisfy.
type Violation struct {
	Rule    string `json:"rule"`
	Message string `json:"message"`
}

// Check parses files (name -> source) and evaluates every rule. A parse
// error is returned as an error; rule failures are returned as violations.
func Check(files map[string][]byte, ruleStrings []string) ([]Violation, error) {
	p := &program{fset: token.NewFileSet()}
	names := make([]string, 0, len(files))
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		f, err := parser.ParseFile(p.fset, n, files[n], parser.SkipObjectResolution)
		if err != nil {
			return nil, err
		}
		p.files = append(p.files, f)
	}
	var out []Violation
	for _, s := range ruleStrings {
		r, err := Parse(s)
		if err != nil {
			return nil, err
		}
		if err := rules[r.Name].check(p, r.Arg); err != nil {
			out = append(out, Violation{Rule: r.String(), Message: err.Error()})
		}
	}
	return out, nil
}

type program struct {
	fset  *token.FileSet
	files []*ast.File
}

func (p *program) funcDecl(name string) *ast.FuncDecl {
	// "Type.Method" selects a method.
	recv, fn, isMethod := strings.Cut(name, ".")
	if !isMethod {
		fn, recv = name, ""
	}
	for _, f := range p.files {
		for _, d := range f.Decls {
			fd, ok := d.(*ast.FuncDecl)
			if !ok || fd.Name.Name != fn {
				continue
			}
			if recv == "" && fd.Recv == nil {
				return fd
			}
			if recv != "" && fd.Recv != nil && len(fd.Recv.List) == 1 && receiverName(fd.Recv.List[0].Type) == recv {
				return fd
			}
		}
	}
	return nil
}

func receiverName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.StarExpr:
		return receiverName(t.X)
	case *ast.Ident:
		return t.Name
	case *ast.IndexExpr: // generic receiver T[P]
		return receiverName(t.X)
	case *ast.IndexListExpr:
		return receiverName(t.X)
	}
	return ""
}

func isSwitch(n ast.Node) bool {
	switch n.(type) {
	case *ast.SwitchStmt, *ast.TypeSwitchStmt:
		return true
	}
	return false
}

func funcContains(what string, match func(ast.Node) bool) func(*program, string) error {
	return func(p *program, fn string) error {
		fd := p.funcDecl(fn)
		if fd == nil || fd.Body == nil {
			return fmt.Errorf("function %s was not found", fn)
		}
		found := false
		ast.Inspect(fd.Body, func(n ast.Node) bool {
			if !found && n != nil && match(n) {
				found = true
			}
			return !found
		})
		if !found {
			return fmt.Errorf("%s should use %s", fn, what)
		}
		return nil
	}
}

func declaresConst(p *program, name string) error {
	for _, f := range p.files {
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok || gd.Tok != token.CONST {
				continue
			}
			for _, s := range gd.Specs {
				for _, id := range s.(*ast.ValueSpec).Names {
					if id.Name == name {
						return nil
					}
				}
			}
		}
	}
	return fmt.Errorf("%s should be declared as a package-level constant", name)
}

// calls checks for a call to pkg.Name, resolving the local name each file
// gives the import (so `import str "strings"` still counts).
func calls(want bool) func(*program, string) error {
	return func(p *program, target string) error {
		pkgPath, fn, _ := strings.Cut(target, ".")
		found := false
		for _, f := range p.files {
			local := localImportName(f, pkgPath)
			if local == "" {
				continue
			}
			ast.Inspect(f, func(n ast.Node) bool {
				call, ok := n.(*ast.CallExpr)
				if !ok {
					return true
				}
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == fn {
					if id, ok := sel.X.(*ast.Ident); ok && id.Name == local {
						found = true
					}
				}
				return !found
			})
		}
		switch {
		case want && !found:
			return fmt.Errorf("the code should call %s", target)
		case !want && found:
			return fmt.Errorf("the code should not call %s; the exercise asks you to do it yourself", target)
		}
		return nil
	}
}

func localImportName(f *ast.File, pkgPath string) string {
	for _, imp := range f.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil || path != pkgPath {
			continue
		}
		if imp.Name != nil {
			if imp.Name.Name == "_" || imp.Name.Name == "." {
				return ""
			}
			return imp.Name.Name
		}
		return path[strings.LastIndex(path, "/")+1:]
	}
	return ""
}

func chanParamsDirectional(p *program, fn string) error {
	fd := p.funcDecl(fn)
	if fd == nil {
		return fmt.Errorf("function %s was not found", fn)
	}
	found := false
	for _, field := range fd.Type.Params.List {
		ct, ok := field.Type.(*ast.ChanType)
		if !ok {
			continue
		}
		found = true
		if ct.Dir == ast.SEND|ast.RECV {
			return fmt.Errorf("channel parameters of %s should be send-only (chan<- T) or receive-only (<-chan T)", fn)
		}
	}
	if !found {
		return fmt.Errorf("%s has no channel parameters", fn)
	}
	return nil
}
