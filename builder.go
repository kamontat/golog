package golog

// Global logger with default setting
var Global *Logger = &Logger{
	level:   INFO,
	setting: &DefaultSetting,
}
