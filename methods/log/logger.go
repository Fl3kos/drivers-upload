package log

import (
	"drivers-create/consts"
	"fmt"
	"log"
	"os"
)

var (
	DebugLog   *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
	initLog    *log.Logger
)

func InitLogger() {
	fileName := fmt.Sprint(consts.LogsRoute)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DebugLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog = log.New(file, "INIT: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog.Println("Init the process")
}

func InitTestLogger() {
	fileName := fmt.Sprint(consts.LogsTestRoute)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DebugLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog = log.New(file, "INIT: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog.Println("Init the process")
}
