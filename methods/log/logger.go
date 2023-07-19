package log

import (
	"fmt"
	"log"
	"os"
	"support-utils/consts"
)

type any = interface{}

var (
	DebugLog   *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
	testLog    *log.Logger
	initLog    *log.Logger
)

func InitLogger() {
	fileName := fmt.Sprint(consts.LogsRoute)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DebugLog = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog = log.New(file, "INIT: ", log.Ldate|log.Ltime|log.Lshortfile)
	testLog = log.New(file, "TEST:", log.Ldate|log.Ltime|log.Lshortfile)
	initLog.Println("Init the process")
}

func InitTestLogger(testName string) {
	fileName := fmt.Sprint(consts.LogsTestRoute)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DebugLog = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog = log.New(file, "TEST INIT: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog.Printf("Init test: %v", testName)
}

func Debugln(text ...any) {
	DebugLog.Println(text...)
}

func Debugf(text string, v ...any) {
	DebugLog.Printf(text, v...)
}

func Warningln(text ...any) {
	WarningLog.Println(text...)
}

func Warningf(text string, v ...any) {
	WarningLog.Printf(text, v...)
}

func Errorln(text ...any) {
	ErrorLog.Println(text...)
}

func Errorf(text string, v ...any) {
	ErrorLog.Printf(text, v...)
}

func Fatalln(text ...any) {
	ErrorLog.Println(text...)
}

func Fatalf(text string, v ...any) {
	ErrorLog.Fatalf(text, v...)
}

func Testln(text ...any) {
	testLog.Println(text...)
}

func Testf(text string, v ...any) {
	testLog.Printf(text, v...)
}
