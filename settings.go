package golog

import (
	"io"
	"os"
)

// Setting of logger struct
type Setting struct {
	withTime bool
	stdout   io.Writer
	stderr   io.Writer
}

// SetTime to true will cause logger to log time level as well
func (s Setting) SetTime(enabled bool) Setting {
	s.withTime = enabled
	return s
}

// SetStdout will set when log write to
func (s Setting) SetStdout(stdout io.Writer) Setting {
	s.stdout = stdout
	return s
}

// SetStderr will set when error write to
func (s Setting) SetStderr(stderr io.Writer) Setting {
	s.stderr = stderr
	return s
}

// DefaultSetting is default setting
var DefaultSetting = Setting{
	withTime: false,
	stdout:   os.Stdout,
	stderr:   os.Stderr,
}
