package golog

import "fmt"

// Logger is logging object
type Logger struct {
	level   Level
	setting *Setting
}

// SetLevel will update current log to specify level
func (l *Logger) SetLevel(level Level) {
	if level == TIME {
		l.SetSetting(l.setting.SetTime(true))
	} else {
		l.level = level
	}
}

// SetSetting will replace setting of logger
func (l *Logger) SetSetting(setting Setting) {
	l.setting = &setting
}

// UpdateSetting will use old setting to create new one
func (l *Logger) UpdateSetting(fn func(s Setting) Setting) {
	s := fn(*l.setting)

	l.SetSetting(s)
}

// IsDebug will create when logging enabled debug mode
func (l *Logger) IsDebug() bool {
	return l.level == DEBUG
}

// IsTime will create when logging enabled timing mode
func (l *Logger) IsTime() bool {
	return l.level == TIME
}

// Stdout will logging data without any formatting
func (l *Logger) Stdout(msg string, params ...interface{}) {
	if l.level != SILENT {
		fmt.Fprintf(l.setting.stdout, msg, params...)
	}
}

// Stderr will logging error without any formatting
func (l *Logger) Stderr(msg string, params ...interface{}) {
	if l.level != SILENT {
		fmt.Fprintf(l.setting.stderr, msg, params...)
	}
}

// CheckLevel will check input level should be log or not
func (l *Logger) CheckLevel(level Level) bool {
	// different condition for time level
	if level == TIME {
		return l.setting.withTime
	} else if level.Code <= l.level.Code {
		return true
	}

	return false
}

func (l *Logger) private(level Level, key int, format string, params ...interface{}) {
	if l.CheckLevel(level) {
		fullFormat := "%s [%03d] " + format

		newParams := make([]interface{}, 0)
		newParams = append(newParams, level.Short)
		newParams = append(newParams, key%1000)
		for _, p := range params {
			newParams = append(newParams, p)
		}

		if level == ERROR || level == WARN {
			l.Stderr(fullFormat, newParams...)
		} else {
			l.Stdout(fullFormat, newParams...)
		}
	}
}

// Newline will go to newline
func (l *Logger) Newline() {
	l.Stdout("\n")
}

// Debug is logging as debug message
func (l *Logger) Debug(key int, format string, params ...interface{}) {
	l.private(DEBUG, key, format, params...)
}

// Time is logging as debug message with time key
func (l *Logger) Time(key int, name string, duration string) {
	l.private(TIME, key, "%-35s: %s", name, duration)
}

// Info is logging as info message
func (l *Logger) Info(key int, format string, params ...interface{}) {
	l.private(INFO, key, format, params...)
}

// Warn is logging as warn message
func (l *Logger) Warn(key int, format string, params ...interface{}) {
	l.private(WARN, key, format, params...)
}

// Error is logging as error message
func (l *Logger) Error(key int, err error) {
	l.private(ERROR, key, err.Error())
}
