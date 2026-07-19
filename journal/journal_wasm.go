//go:build js

package journal

func Enabled() bool { return false }

func Send(message string, priority Priority, vars map[string]string) error { return nil }
