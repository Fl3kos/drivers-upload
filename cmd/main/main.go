package main

import (
	convert "drivers-create/methods/converts"
	csv "drivers-create/methods/csv"
	dniM "drivers-create/methods/dni"
	files "drivers-create/methods/file"
	json "drivers-create/methods/json"
	logs "drivers-create/methods/log"
	sql "drivers-create/methods/sql"
	"fmt"
	"strings"
)

func main() {
	logs.InitLogger()

	allDnis := getAllDnis()

	dnisIncorrect, err := dniM.ComprobeAllDnis(allDnis)
	if err == nil {
		allUsers := convert.ConvertAllDnisToUsers(allDnis)
		allPasswords := convert.ConvertAllUsersToPasswords(allUsers)

		allNames := getAllNames()
		allPhones := getAllPhones()

		shops := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
		shops = shops[:len(shops)-1]

		shopCodes, shopNames := getShopCodesAndShopNames(shops)

		exportDriversToCsv(allUsers, allNames, allPasswords, shopNames)

		// creacion de las queries
		jsonT := json.GenerateJson(allNames, allPasswords, allUsers)
		namesT := convert.TransformAllNames(allNames)
		driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)
		relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCodes)
		sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

		// creacion de los files
		files.GenerateFile(jsonT, files.CreationFileRoute("usersCouchbase", "json"))
		files.GenerateFile(namesT, files.CreationFileRoute("names", "txt"))
		files.GenerateFile(sqlLiteInserts, files.CreationFileRoute("insertSQLIteQuery", "sql"))

	} else {
		logs.ErrorLog.Printf("Error validating dnis, incorrect DNIs: %v. Error %v", dnisIncorrect, err)
		fmt.Println("Error without validate dnis, check the logs")
	}

	fmt.Println("Finish")
}

//metodo que devuelva todos los dnis
func getAllDnis() []string {
	dnis := files.ReadFile(files.ReadFileRoute("dnis", "txt"))
	dnis = strings.ToUpper(dnis)

	substring := dnis[:len(dnis)-1]
	allDnis := strings.Split(substring, "\n")

	//make trim to all dnis
	for _, dni := range allDnis {
		if dni != "" {
			dni = strings.TrimSpace(dni)
		}

	}

	return allDnis
}

func getAllNames() []string {
	names := files.ReadFile(files.ReadFileRoute("names", "txt"))
	allNames := strings.Split(names, "\n")

	//make trim to all users
	for i, name := range allNames {
		if name != "" {
			allNames[i] = strings.TrimSpace(name)
		}

	}

	return allNames
}

func getAllPhones() []string {
	phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
	allPhones := strings.Split(phonesNumber, "\n")

	return allPhones
}

func getShopCodesAndShopNames(shops []string) ([]string, []string) {
	var shopCodes []string
	var shopNames []string

	for i, _ := range shops {
		shop := strings.Split(shops[i], "-")

		shopCode := shop[0]
		shopName := shop[1]

		shopCode = strings.TrimSpace(shopCode)
		shopName = strings.TrimSpace(shopName)

		shopCodes = append(shopCodes, shopCode)
		shopNames = append(shopNames, shopName)
	}

	return shopCodes, shopNames
}

func exportDriversToCsv(allUsers, allNames, allPasswords, shopsNames []string) {
	nextShop := 0

	for i := 0; i < len(shopsNames); i++ {
		drivers := []csv.Driver{}
		for j := nextShop; j < len(allNames); j++ {
			if allNames[j] == "" {
				nextShop = j + 1
				j = len(allNames)
				break
			}
			driver := getDriver(allNames[j], allUsers[j], allPasswords[j])
			drivers = append(drivers, driver)
		}
		csv.ExportCsvFile(drivers, shopsNames[i])
	}

}

func getDriver(name, user, password string) csv.Driver {
	driver := csv.Driver{}

	driver = csv.Driver{name, user, password}

	return driver
}
