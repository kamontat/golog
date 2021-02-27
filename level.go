package golog

import "strings"

// Level for logging
type Level struct {
	Code  uint8
	Name  string
	Short string
}

func (l Level) String() string {
	return strings.ToLower(l.Name)
}
