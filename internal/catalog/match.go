package catalog

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// NotFoundError reports an unknown exercise query with nearby suggestions.
type NotFoundError struct {
	Query       string
	Suggestions []string
}

func (e *NotFoundError) Error() string {
	msg := fmt.Sprintf("no exercise matches %q", e.Query)
	if len(e.Suggestions) > 0 {
		msg += "; did you mean: " + strings.Join(e.Suggestions, ", ") + "?"
	}
	return msg
}

// AmbiguousError reports a prefix that matches more than one exercise.
type AmbiguousError struct {
	Query   string
	Matches []string
}

func (e *AmbiguousError) Error() string {
	return fmt.Sprintf("%q is ambiguous; it matches: %s", e.Query, strings.Join(e.Matches, ", "))
}

// Find resolves a user query to an exercise. It accepts an exact ID, a
// display ordinal ("3"), or an unambiguous ID prefix ("vari"). It never
// silently picks among several matches.
func (c *Catalog) Find(query string) (*Exercise, error) {
	q := strings.ToLower(strings.TrimSpace(query))
	if ex, ok := c.byID[q]; ok {
		return ex, nil
	}
	if n, err := strconv.Atoi(q); err == nil {
		if n >= 1 && n <= len(c.exercises) {
			return c.exercises[n-1], nil
		}
		return nil, &NotFoundError{Query: query}
	}
	var matches []string
	for _, ex := range c.exercises {
		if strings.HasPrefix(ex.ID, q) {
			matches = append(matches, ex.ID)
		}
	}
	switch len(matches) {
	case 1:
		return c.byID[matches[0]], nil
	case 0:
		return nil, &NotFoundError{Query: query, Suggestions: c.suggest(q)}
	default:
		return nil, &AmbiguousError{Query: query, Matches: matches}
	}
}

func (c *Catalog) suggest(q string) []string {
	type cand struct {
		id   string
		dist int
	}
	var cands []cand
	for _, ex := range c.exercises {
		d := levenshtein(q, ex.ID)
		// Also compare against the topic part, so "slice" suggests slices_01.
		if td := levenshtein(q, ex.Topic); td < d {
			d = td
		}
		if d <= 3 || (len(q) >= 3 && strings.Contains(ex.ID, q)) {
			cands = append(cands, cand{ex.ID, d})
		}
	}
	sort.SliceStable(cands, func(i, j int) bool { return cands[i].dist < cands[j].dist })
	var out []string
	for i := 0; i < len(cands) && i < 5; i++ {
		out = append(out, cands[i].id)
	}
	return out
}

func levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	cur := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev, cur = cur, prev
	}
	return prev[len(rb)]
}
