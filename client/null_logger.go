package client

type NullLogger struct {
}

func (n NullLogger) Debug(string)   {}
func (n NullLogger) Crit(string)    {}
func (n NullLogger) Info(string)    {}
func (n NullLogger) Warning(string) {}
func (n NullLogger) Err(string)     {}
func (n NullLogger) Emerg(string)   {}
