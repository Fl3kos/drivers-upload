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
		drivers := getDrivers(allUsers, allNames, allPasswords)
		shopCode := strings.Split(files.ReadFile(files.ReadFileRoute("shopCode", "txt")), "\n")[0]

		// creacion de las queries
		jsonT := json.GenerateJson(allNames, allPasswords, allUsers)
		namesT := convert.TransformAllNames(allNames)
		driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)
		relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCode)
		sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

		// creacion de los files
		files.GenerateFile(jsonT, files.CreationFileRoute("usersCouchbase", "json"))
		files.GenerateFile(namesT, files.CreationFileRoute("names", "txt"))
		files.GenerateFile(sqlLiteInserts, files.CreationFileRoute("insertSQLIteQuery", "sql"))
		csv.ExportCsvFile(drivers)
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
	for i, dni := range allDnis {
		allDnis[i] = strings.TrimSpace(dni)
	}

	return allDnis
}

func getAllNames() []string {
	names := files.ReadFile(files.ReadFileRoute("names", "txt"))
	allNames := strings.Split(names, "\n")

	//make trim to all users
	for i, name := range allNames {
		allNames[i] = strings.TrimSpace(name)
	}

	return allNames
}

func getAllPhones() []string {
	phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
	allPhones := strings.Split(phonesNumber, "\n")

	return allPhones
}

func getDrivers(allUsers, allNames, allPasswords []string) []csv.Driver {
	drivers := []csv.Driver{}

	for i, _ := range allUsers {
		driver := csv.Driver{allNames[i], allUsers[i], allPasswords[i]}
		drivers = append(drivers, driver)
	}

	return drivers
}
