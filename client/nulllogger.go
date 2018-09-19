package client

// NullLogger don't log anything when using as lib
type NullLogger struct {
}

// Debug noop
func (n NullLogger) Debug(string) {}

// Crit noop
func (n NullLogger) Crit(string) {}

// Info noop
func (n NullLogger) Info(string) {}

// Warning noop
func (n NullLogger) Warning(string) {}

// Err noop
func (n NullLogger) Err(string) {}

// Emerg noop
func (n NullLogger) Emerg(string) {}
