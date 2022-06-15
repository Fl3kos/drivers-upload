package main

import (
	convert "drivers-create/methods/converts"
	csv "drivers-create/methods/csv"
	dniM "drivers-create/methods/dni"
	files "drivers-create/methods/file"
	json "drivers-create/methods/json"
	sql "drivers-create/methods/sql"
	"strings"
)

func main() {
	//readfiles method

	dnis := files.ReadFile(files.ReadFileRoute("dnis", "txt"))
	dnis = strings.ToUpper(dnis)

	substring := dnis[:len(dnis)-1]
	allDnis := strings.Split(substring, "\n")

	//make trim to all dnis
	for i, dni := range allDnis {
		allDnis[i] = strings.TrimSpace(dni)
	}

	_continue := dniM.ComprobeAllDnis(allDnis)
	//_continue := comprobeAllDnis()

	// convertir los dni en usuarios
	if _continue {
		allUsers := convert.ConvertAllDnisToUsers(allDnis)

		// convertir las contraseñas si no existen
		allPasswords := convert.ConvertAllUsersToPasswords(allUsers)

		//fmt.Println("pon los nombres separdos por comas (,)")
		names := files.ReadFile(files.ReadFileRoute("names", "txt"))
		allNames := strings.Split(names, "\n")

		phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
		allPhones := strings.Split(phonesNumber, "\n")

		shopCode := strings.Split(files.ReadFile(files.ReadFileRoute("shopCode", "txt")), "\n")[0]

		//make trim to all users
		for i, name := range allNames {
			allNames[i] = strings.TrimSpace(name)
		}

		// creacion de las queries
		jsonT := json.GenerateJson(allNames, allPasswords, allUsers)
		namesT := convert.TransformAllNames(allNames)
		driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)
		relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCode)
		sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

		//passwords := strings.Join(allPasswords, ",")

		// creacion de los files
		files.GenerateFile(jsonT, files.CreationFileRoute("usersCouchbase", "json"))
		files.GenerateFile(namesT, files.CreationFileRoute("names", "txt"))
		files.GenerateFile(convert.UsersAndPasswords(allNames, allUsers, allPasswords), files.CreationFileRoute("usersAndPasswords", "txt"))
		files.GenerateFile(sqlLiteInserts, files.CreationFileRoute("insertSQLIteQuery", "sql"))

		drivers := []csv.Driver{}

		for i, _ := range allUsers {
			driver := csv.Driver{allNames[i], allUsers[i], allPasswords[i]}
			drivers = append(drivers, driver)
		}

		csv.ExportCsvFile(drivers)
		//WriteCsv()
	}

	// futuramente hacer un menú
}
