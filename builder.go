package golog

// Global logger with default setting
var Global *Logger = &Logger{
	level:   INFO,
	setting: &DefaultSetting,
}

// New will create new logger struct
func New(level Level, setting Setting) *Logger {
	return &Logger{
		level:   level,
		setting: &setting,
	}
}
