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
)

func main() {
	logs.InitLogger()

	allDnis := getAllDnis()

	dnisIncorrect, err := dniM.ComprobeAllDnis(allDnis)
	if err == nil {
		allUsers := dniToUser.ConvertAllDnisToUsers(allDnis)
		allPasswords := userToPassword.ConvertAllUsersToPasswords(allUsers)

		allNames := getAllNames()
		allPhones := getAllPhones()

		shops := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
		shops = shops[:len(shops)-1]

		shopCodes, shopNames := getShopCodesAndShopNames(shops)

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

		//change white space to -
		shopName = strings.ReplaceAll(shopName, " ", "_")
		shopCodes = append(shopCodes, shopCode)
		shopNames = append(shopNames, shopName)
	}

	return shopCodes, shopNames
}
