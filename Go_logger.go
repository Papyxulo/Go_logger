package Go_logger

import (
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	Level int
}

func (logger Logger) check_verbosity_level(local_level int) bool {
	return logger.Level >= local_level
}

func (logger Logger) Fatal(text string, code int) {
	if !logger.check_verbosity_level(1) {
		return
	}
	var fatal = color.New(color.FgRed)
	fatal = fatal.Add(color.Bold)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fatal.Println(formated_time + " - " + text + " - Run the script on verbose level 3 to get debug info")
	os.Exit(code)
}

func (logger Logger) Info(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}
	info := color.New(color.FgBlue)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	info.Println(formated_time + " - " + text)
}

func (logger Logger) Sucess(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}
	Sucess := color.New(color.FgGreen)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	Sucess.Println(formated_time + " - " + text)
}

func (logger Logger) Failed(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}
	var failed = color.New(color.FgRed)
	failed = failed.Add(color.Bold)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	failed.Println(formated_time + " - " + text)
}

func (logger Logger) Error(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}
	var failed = color.New(color.FgRed)
	failed = failed.Add(color.Bold)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	failed.Println(formated_time + " - ERROR: " + text)
}

func (logger Logger) Trace(text string) {
	if !logger.check_verbosity_level(2) {
		return
	}
	var failed = color.New(color.FgRed)
	failed = failed.Add(color.Bold)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	failed.Println(formated_time + " - " + text)
}

func (logger Logger) Warning(text string) {
	if !logger.check_verbosity_level(2) {
		return
	}
	var warning = color.New(color.FgYellow)
	warning = warning.Add(color.Bold)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	warning.Println(formated_time + " - " + text)
}

func (logger Logger) Debug(text string) {
	if !logger.check_verbosity_level(3) {
		return
	}
	Debug := color.New(color.FgMagenta)

	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	Debug.Println(formated_time + " - " + text)
}
