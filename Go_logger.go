package Go_logger

import (
	"fmt"
	"os"
	"runtime"
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

	color.Set(color.FgRed, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v - Run the script on verbose level 3 to get debug info \n",
		formated_time, text)
	color.Unset()
	os.Exit(code)
}

func (logger Logger) Info(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}

	color.Set(color.FgBlue)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()
}

func (logger Logger) Sucess(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}

	color.Set(color.FgGreen)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()
}

func (logger Logger) Failed(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}

	color.Set(color.FgRed, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()
}

func (logger Logger) Error(text string) {
	if !logger.check_verbosity_level(1) {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	function_name := string(runtime.FuncForPC(pc).Name())

	color.Set(color.FgRed, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v ERROR: %v\n", formated_time, function_name, text)

	color.Unset()
}

func (logger Logger) Trace(text string, objs ...interface{}) {
	if !logger.check_verbosity_level(2) {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	function_name := string(runtime.FuncForPC(pc).Name())

	color.Set(color.FgYellow, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v - %v\n", formated_time, function_name, text)

	// if there are object to print
	for _, obj := range objs {
		fmt.Printf("\t %v\n", obj)
	}

	color.Unset()
}

func (logger Logger) Warning(text string) {
	if !logger.check_verbosity_level(2) {
		return
	}

	color.Set(color.FgYellow, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()
}

func (logger Logger) Debug(text string, objs ...interface{}) {
	if !logger.check_verbosity_level(3) {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	function_name := string(runtime.FuncForPC(pc).Name())

	color.Set(color.FgMagenta)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v - %v\n", formated_time, function_name, text)

	// if there are object to print
	for _, obj := range objs {
		fmt.Printf("\t %v\n", obj)
	}

	color.Unset()
}
