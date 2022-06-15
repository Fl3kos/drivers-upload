package main

import (
	convert "drivers-create/methods/converts"
	dniM "drivers-create/methods/dni"
	files "drivers-create/methods/file"
	json "drivers-create/methods/json"
	sql "drivers-create/methods/sql"
	"strings"
)

//var allDnis []string
var allUsers []string
var allPasswords []string
var allNames []string
var allPhones []string
var shopCode string

var m int

func main() {
	//readfiles method

	dnis := files.ReadFile(files.ReadFileRoute("dnis", "txt"))

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
		m = len(allDnis)
		allUsers = convert.ConvertAllDnisToUsers(allDnis)

		// convertir las contraseñas si no existen
		allPasswords = convert.ConvertAllUsersToPasswords(allUsers)

		//fmt.Println("pon los nombres separdos por comas (,)")
		names := files.ReadFile(files.ReadFileRoute("names", "txt"))
		allNames = strings.Split(names, "\n")

		phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
		allPhones = strings.Split(phonesNumber, "\n")

		shopCode = strings.Split(files.ReadFile(files.ReadFileRoute("shopCode", "txt")), "\n")[0]

		//make trim to all users
		for i, name := range allNames {
			allNames[i] = strings.TrimSpace(name)
		}

		// creacion de las queries
		jsonT := json.GenerateJson(allNames, allPasswords, allUsers, m)
		sqlT := sql.GenerateSql(allUsers, m)
		namesT := convert.TransformAllNames(allNames)
		driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones, m)
		relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCode, m)
		sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

		//passwords := strings.Join(allPasswords, ",")

		// creacion de los files
		files.GenerateFile(jsonT, files.CreationFileRoute("usersCouchbase", "json"))
		files.GenerateFile(sqlT, files.CreationFileRoute("insertQuery", "sql"))
		files.GenerateFile(namesT, files.CreationFileRoute("names", "txt"))
		files.GenerateFile(convert.UsersAndPasswords(allNames, allUsers, allPasswords, m), files.CreationFileRoute("usersAndPasswords", "txt"))
		files.GenerateFile(sqlLiteInserts, files.CreationFileRoute("insertSQLIteQuery", "sql"))
		//WriteCsv()
		// end
	}

	// futuramente hacer un menú
}

//this method generate json text
//the length of arrays is the value of x
