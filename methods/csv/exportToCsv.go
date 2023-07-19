package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	files "support-utils/methods/file"
	logs "support-utils/methods/log"
)

type Driver struct {
	Name string
	User string
	Pass string
}

func ExportDriversToCsv(allUsers, allNames, allPasswords, shopsNames []string) {
	nextShop := 0

	for i := 0; i < len(shopsNames); i++ {
		drivers := []Driver{}
		for j := nextShop; j < len(allNames); j++ {
			if allNames[j] == "" {
				nextShop = j + 1
				j = len(allNames)
				break
			}
			driver := getDriver(allNames[j], allUsers[j], allPasswords[j])
			drivers = append(drivers, driver)
		}
		exportCsvFile(drivers, shopsNames[i])
	}

}

func getDriver(name, user, password string) Driver {
	driver := Driver{}

	driver = Driver{name, user, password}

	return driver
}

func exportCsvFile(drivers []Driver, shopName string) {
	logs.Debugln("Exporting Csv file with names, users and passwords")

	fileName := fmt.Sprintf("userAndPassword-%v", shopName)
	file, err := os.Create(files.CreationFileRouteCsv(fileName, "csv"))

	defer file.Close()

	if err != nil {
		logs.Errorln("Failed to open file", err)
	}

	w := csv.NewWriter(file)

	defer w.Flush()

	// Using Write
	row := []string{"Name", "User", "Password"}
	w.Write(row)

	for _, record := range drivers {
		row = []string{record.Name, record.User, record.Pass}
		if err := w.Write(row); err != nil {
			logs.Errorln("error writing record to file", err)
		} else {
			logs.Debugln("Users and passwords file is writig")
		}
	}
}
