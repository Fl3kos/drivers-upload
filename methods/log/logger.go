package log

import (
	"fmt"
	"log"
	"os"

	common "drivers-create/methods"
)

var (
	DebugLog   *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
	initLog    *log.Logger
)

func InitLogger() {
	fileName := fmt.Sprint("./logs/log-", common.GetDate(), ".txt")
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
