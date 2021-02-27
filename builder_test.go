package golog_test

import (
	"bytes"
	"testing"

	"github.com/kamontat/golog"
)

type W struct {
	Data string
}

func (w *W) Write(data []byte) (int, error) {
	w.Data = bytes.NewBuffer(data).String()

	return -1, nil
}

func TestGlobalLogger(t *testing.T) {
	debugMode := golog.Global.IsDebug()
	if debugMode {
		t.Error("Default setting should disable debug be default")
	}
}

func TestLogInfo(t *testing.T) {
	writer := &W{}
	golog.Global.SetSetting(golog.DefaultSetting.SetStdout(writer))
	golog.Global.Info(1, "print")

	if writer.Data == "" {
		t.Error("Data didn't write to io.writer")
	}
}

func TestLogDebug(t *testing.T) {
	writer := &W{}
	golog.Global.SetSetting(golog.DefaultSetting.SetStdout(writer))
	golog.Global.Debug(1, "print")

	if writer.Data != "" {
		t.Error("Data shouldn't be log")
	}
}
