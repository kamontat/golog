package golog_test

import (
	"testing"

	"github.com/kamontat/golog"
)

func TestDefaultValue(t *testing.T) {
	if &golog.DefaultSetting == nil {
		t.Error("Default setting cannot be nil")
	}
}

func TestUpdateDefaultValue(t *testing.T) {
	a := golog.DefaultSetting
	b := golog.DefaultSetting.SetTime(true)

	if a == b {
		t.Error("When update setting, it should return new object")
	}
}
