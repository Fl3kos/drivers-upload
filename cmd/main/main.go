package main

import (
	"drivers-create/consts"
	convert "drivers-create/methods/converts"
	csv "drivers-create/methods/csv"
	dniM "drivers-create/methods/dni"
	"drivers-create/methods/dniToUser"
	files "drivers-create/methods/file"
	json "drivers-create/methods/json"
	logs "drivers-create/methods/log"
	sql "drivers-create/methods/sql"
	sqlite "drivers-create/methods/sqlite"
	"drivers-create/methods/userToPassword"
	"fmt"
	"strings"
	"drivers-create/methods/getDnis"
	"drivers-create/methods/getNames"
	"drivers-create/methods/getPhones"
	"drivers-create/methods/getShops"
)

func main() {
	logs.InitLogger()

	allDnis := getDnis.GetAllDnis()

	dnisIncorrect, err := dniM.ComprobeAllDnis(allDnis)
	if err == nil {
		allUsers := dniToUser.ConvertAllDnisToUsers(allDnis)
		allPasswords := userToPassword.ConvertAllUsersToPasswords(allUsers)

		allNames := getNames.GetAllNames()
		allPhones := getPhones.GetAllPhones()

		shops := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
		shops = shops[:len(shops)-1]

		shopCodes, shopNames := getShops.GetShopCodesAndShopNames(shops)

		csv.ExportDriversToCsv(allUsers, allNames, allPasswords, shopNames)

		// queries creation
		jsonT := json.GenerateJson(allNames, allPasswords, allUsers)
		namesT := convert.TransformAllNames(allNames)
		driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)
		relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCodes)
		sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

		//insert in sqlite
		err := sqlite.InsertSqlite(sqlLiteInserts, consts.SqliteDatabase)
		if err != nil {
			fmt.Println("Error inserting drivers in database. Check the logs")
			logs.ErrorLog.Printf("Error insert in database. Error: %v", err)
		}

		// files created
		err = files.GenerateFile(jsonT, files.CreationFileRouteJson("usersCouchbase", "json"))
		if err != nil {
			logs.ErrorLog.Printf("Error generating file: %v", err)
			fmt.Println("Error generating files, check the logs /logs/lo")
		}

		err = files.GenerateFile(namesT, files.CreationFileRouteNames("names", "txt"))
		if err != nil {
			logs.ErrorLog.Printf("Error generating file: %v", err)
			fmt.Println("Error generating files, check the logs /logs/lo")
		}

		err = files.GenerateFile(sqlLiteInserts, files.CreationFileRouteSql("insertSQLIteQuery", "sql"))
		if err != nil {
			logs.ErrorLog.Printf("Error generating file: %v", err)
			fmt.Println("Error generating files, check the logs /logs/lo")
		}

	} else {
		logs.ErrorLog.Printf("Error validating dnis, incorrect DNIs: %v. Error %v", dnisIncorrect, err)
		fmt.Println("Error without validate dnis, check the logs")
	}

	fmt.Println("Finish")
}
