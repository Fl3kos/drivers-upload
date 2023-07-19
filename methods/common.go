package methods

import (
	"fmt"
	"strconv"
	"strings"
	logs "support-utils/methods/log"
	"time"
)

func IsNumber(c string) bool {
	_, err := strconv.Atoi(c)
	if err == nil {
		return true
	}
	return false
}

func exportToDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func GetDate() string {
	var lines []string
	var toDate string
	var today = exportToDate()
	for _, line := range strings.Split(today, "-") {
		line = strings.TrimSpace(line)

		if line != "" {
			lines = append(lines, line)
		}

		toDate = toDate + line
	}

	return toDate
}

func ControlErrors(err error) {
	if err != nil {
		logs.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/lo")
	}
}
