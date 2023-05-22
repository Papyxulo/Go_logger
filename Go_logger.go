package Go_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	Level   int
	History []string
}

func (logger Logger) CheckVerbosityLevel(local_level int) bool {
	return logger.Level >= local_level
}

func (logger *Logger) Fatal(text string, code int) {
	if !logger.CheckVerbosityLevel(1) {
		return
	}

	color.Set(color.FgRed, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v - Run the script on verbose level 3 to get debug info \n",
		formated_time, text)
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v", formated_time, text))
	os.Exit(code)
}

func (logger *Logger) Info(text string) {
	if !logger.CheckVerbosityLevel(1) {
		return
	}

	color.Set(color.FgBlue)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v", formated_time, text))
}

func (logger *Logger) Sucess(text string) {
	if !logger.CheckVerbosityLevel(1) {
		return
	}

	color.Set(color.FgGreen)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v", formated_time, text))
}

func (logger *Logger) Failed(text string) {
	if !logger.CheckVerbosityLevel(1) {
		return
	}

	color.Set(color.FgRed, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v", formated_time, text))
}

func (logger *Logger) Error(text string) {
	if !logger.CheckVerbosityLevel(1) {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	function_name := string(runtime.FuncForPC(pc).Name())

	color.Set(color.FgRed, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v ERROR: %v\n", formated_time, function_name, text)
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v", formated_time, text))
}

func (logger *Logger) Trace(text string, objs ...interface{}) {
	if !logger.CheckVerbosityLevel(2) {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	function_name := string(runtime.FuncForPC(pc).Name())

	color.Set(color.FgYellow, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v - %v\n", formated_time, function_name, text)

	// if there are object to print
	for _, obj := range objs {
		fmt.Printf("\t %+v\n", obj)
	}
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v - %v", formated_time, function_name, text))
}

func (logger *Logger) Warning(text string) {
	if !logger.CheckVerbosityLevel(2) {
		return
	}

	color.Set(color.FgYellow, color.Bold)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v\n", formated_time, text)
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v", formated_time, text))
}

func (logger *Logger) Debug(text string, objs ...interface{}) {
	if !logger.CheckVerbosityLevel(3) {
		return
	}

	pc, _, _, _ := runtime.Caller(1)
	function_name := string(runtime.FuncForPC(pc).Name())

	color.Set(color.FgMagenta)
	var formated_time = string(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v - %v - %v\n", formated_time, function_name, text)

	// if there are object to print
	for _, obj := range objs {
		fmt.Printf("\t %+v\n", obj)
	}
	color.Unset()

	//logs for panic
	logger.AppendHistory(fmt.Sprintf("%v - %v - %v", formated_time, function_name, text))

}

func (logger Logger) PanicHandler() {
	logs := logger.GetHistory()

	r := recover()
	if r != nil {
		fmt.Println("\n----------PANIC-------------\n", r)
		fmt.Println("Panic caused by error :", r)
		fmt.Printf("Logs :\n\t%v\n", logs)
		fmt.Println("\n----------STACK-------------\n", r)
		stack := string(debug.Stack())
		fmt.Printf("%v", stack)

		// get current location
		ex, err := os.Executable()
		if err != nil {
			fmt.Printf("Error generating panic file : %v", err.Error())
		}
		// creating file
		file := filepath.Dir(ex) + "/panic.txt"
		data :=
			fmt.Sprintf("--------------------------\npanic caused by error : %v\nLogs : \n\t%v\n"+
				"Stack : \n%v\n-----------------------", r, logs, stack)

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("Error opening panic file for writing : %v", err.Error())
		}
		defer f.Close()

		_, err2 := f.WriteString(data)
		if err2 != nil {
			fmt.Printf("Error writing data to panic file : %v", err.Error())
		}

	}
}

func (logger *Logger) AppendHistory(text string) {
	(*logger).History = append((*logger).History, text)
}

func (logger Logger) GetHistory() []string {
	return logger.History
}
