package logs

func (s *Spitter) Error(a string) {
	s.spitLog(LevelError, a)
}
func (s *Spitter) Warning(a string) {
	s.spitLog(LevelWarning, a)
}

func (s *Spitter) Info(a string) {
	s.spitLog(LevelInfo, a)
}

func (s *Spitter) Debug(a string) {
	s.spitLog(LevelDebug, a)
}
