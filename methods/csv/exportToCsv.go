package csv

import (
	files "drivers-create/methods/file"
	logs "drivers-create/methods/log"
	"encoding/csv"
	"fmt"
	"os"
)

type Driver struct {
	Name string
	User string
	Pass string
}

func ExportCsvFile(drivers []Driver, shopName string) {
	logs.DebugLog.Println("Exporting Csv file with names, users and passwords")

	fileName := fmt.Sprintf("userAndPassword-%v", shopName)
	file, err := os.Create(files.CreationFileRoute(fileName, "csv"))

	defer file.Close()

	if err != nil {
		logs.ErrorLog.Println("Failed to open file", err)
	}

	w := csv.NewWriter(file)

	defer w.Flush()

	// Using Write
	row := []string{"Name", "User", "Password"}
	w.Write(row)

	for _, record := range drivers {
		row = []string{record.Name, record.User, record.Pass}
		if err := w.Write(row); err != nil {
			logs.ErrorLog.Println("error writing record to file", err)
		} else {
			logs.DebugLog.Println("Users and passwords file is writig")
		}
	}
}
