package log

import (
	"github.com/labstack/gommon/log"
	"io"
)

/*
const (
	PRINT Lvl = iota + 1
	DEBUG
	INFO
	WARN
	ERROR
	PANIC
	FATAL
	OFF
)
*/

type (
	Lvl    uint8
	JSON   interface{}
	Logger interface {
		Level() Lvl
		SetLevel(v Lvl)

		Output() io.Writer
		SetOutput(w io.Writer)

		Prefix() string
		SetPrefix(prefix string)

		Flags() int
		SetFlags(flag int)

		SetHeader(h string)

		Printf(format string, v ...interface{})
		Print(v ...interface{})
		Println(v ...interface{})
		PrintJSON(v JSON)

		Debug(v ...interface{})
		Debugf(format string, v ...interface{})
		Debugln(v ...interface{})
		DebugJSON(v JSON)

		Info(v ...interface{})
		Infof(format string, v ...interface{})
		Infoln(v ...interface{})
		InfoJSON(v log.JSON)

		Warn(i ...interface{})
		Warnf(format string, args ...interface{})
		Warnln(v ...interface{})
		WarnJSON(v log.JSON)

		Error(i ...interface{})
		Errorf(format string, args ...interface{})
		Errorj(j log.JSON)
		ErrorJSON(v log.JSON)

		Fatal(v ...interface{})
		Fatalf(format string, v ...interface{})
		Fatalln(v ...interface{})
		FatalJSON(v JSON)

		Panic(v ...interface{})
		Panicf(format string, v ...interface{})
		Panicln(v ...interface{})
		PanicJSON(v JSON)
	}
)
