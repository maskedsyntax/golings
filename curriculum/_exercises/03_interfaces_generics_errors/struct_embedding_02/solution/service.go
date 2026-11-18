// struct_embedding_02: Overriding and delegating (reference solution)
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

// Log records msg prefixed with the service name. It shadows Logger.Log and
// delegates to it through the embedded field.
func (s *Service) Log(msg string) {
	s.Logger.Log("[" + s.Name + "] " + msg)
}
