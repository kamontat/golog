# Golog (kamontat)

logger utilities

## Usage

Installing via `go get github.com/kamontat/golog`.

```golang

import (
	"os"
	"time"

	"github.com/kamontat/golog"
)

func normal() {
	golog.Global.Stdout("hello world")
}

func customSetting() {
	golog.Global.SetSetting(golog.DefaultSetting.SetTime(true))

	dur, err := time.ParseDuration("10h")

	golog.Global.Error(1, err)
	golog.Global.Time(1, "Name", dur.String())
}

func localLogger() {
	logger := golog.New(golog.WARN, golog.DefaultSetting.SetStdout(os.Stderr))

	logger.Warn(1, "hello world")
}

```