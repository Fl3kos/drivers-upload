package csv

import (
	"drivers-create/methods/file"
	"encoding/csv"
	"log"
	"os"
)

type Driver struct {
	Name string
	User string
	Pass string
}

func ExportCsvFile(drivers []Driver) {

	file, err := os.Create(file.CreationFileRoute("userAndPassword", "csv"))
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	// Using Write
	row := []string{"Name", "User", "Password"}
	w.Write(row)
	for _, record := range drivers {
		row = []string{record.Name, record.User, record.Pass}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
