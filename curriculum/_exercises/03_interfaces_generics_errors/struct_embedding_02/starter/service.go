// struct_embedding_02: Overriding and delegating
//
// Give *Service its own Log method. See README.md.
//
//	api := &Service{Logger: l, Name: "api"}
//	api.Log("started") // the logger records "[api] started"
//
// Reference: https://gobyexample.com/struct-embedding
package service

// Logger collects log lines.
type Logger struct {
	lines []string
}

// Log records one line.
func (l *Logger) Log(line string) {
	l.lines = append(l.lines, line)
}

// Lines returns everything logged so far.
func (l *Logger) Lines() []string {
	return l.lines
}

// Service is a named component that logs through a shared Logger.
type Service struct {
	*Logger
	Name string
}

// TODO: func (s *Service) Log(msg string)
