package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if msg == "" {
		fmt.Println("")
		return
	}

	c := color.New(color.FgHiCyan)
	c.Println(fmt.Sprintf(msg, args...))
}

func (l *Logger) Err(err error) {
	c := color.New(color.FgHiRed)
	c.Println(fmt.Sprintf("%#v", err))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	c := color.New(color.FgHiRed)
	c.Println(fmt.Sprintf(msg, args...))

}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Instructions(msg, args...)
}

func (l *Logger) Instructions(msg string, args ...interface{}) {
	white := color.New(color.FgHiWhite)
	//white.Println("")
	white.Println(fmt.Sprintf(msg, args...))
}
