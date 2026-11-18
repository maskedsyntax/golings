// Mutant: forgets the space after the prefix.
package service

type Logger struct {
	lines []string
}

func (l *Logger) Log(line string) { l.lines = append(l.lines, line) }

func (l *Logger) Lines() []string { return l.lines }

type Service struct {
	*Logger
	Name string
}

func (s *Service) Log(msg string) { s.Logger.Log("[" + s.Name + "]" + msg) }
