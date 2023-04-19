package main

import (
	"drivers-create/consts"
	convert "drivers-create/methods/converts"
	csv "drivers-create/methods/csv"
	dniM "drivers-create/methods/dni"
	"drivers-create/methods/dniToUser"
	files "drivers-create/methods/file"
	"drivers-create/methods/gets/getDnis"
	"drivers-create/methods/gets/getNames"
	"drivers-create/methods/gets/getPhones"
	"drivers-create/methods/gets/getShops"
	"drivers-create/methods/http"
	json "drivers-create/methods/json"
	logs "drivers-create/methods/log"
	sql "drivers-create/methods/sql"
	sqlite "drivers-create/methods/sql/sqlite"
	"drivers-create/methods/userToPassword"
	"fmt"
	"strings"
)

func main() {
	logs.InitLogger()

	allDnis := getDnis.GetAllDnis()

	dnisIncorrect, err := dniM.ComprobeAllDnis(allDnis)
	if err != nil {
		fmt.Println("Error without validate dnis, check the logs")
		logs.Fatalf("Error validating dnis, incorrect DNIs: %v. Error: %v", dnisIncorrect, err)
	}

	allUsers := dniToUser.ConvertAllDnisToUsers(allDnis)
	allPasswords := userToPassword.ConvertAllUsersToPasswords(allUsers)

	allNames := getNames.GetAllNames()
	allPhones := getPhones.GetAllPhones()

	shops := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
	shops = shops[:len(shops)-1]

	shopCodes, shopNames := getShops.GetShopCodesAndShopNames(shops)

	csv.ExportDriversToCsv(allUsers, allNames, allPasswords, shopNames)

	// Create Json files
	jsonAcl := json.GenerateJson(allNames, allPasswords, allUsers, allPhones, shopCodes)
	jsonEndPoint := json.GenerateEndpointJson(allNames, allPasswords, allUsers, allPhones, shopCodes)

	// Create Names file
	namesT := convert.TransformAllNames(allNames)

	// Create SQL files
	driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)
	relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCodes)
	sqlAcl := sql.GenerateAclInsert(allUsers, "ROLE_APPTMS_DRIVER")
	sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

	//insert in sqlite
	err = sqlite.InsertSqlite(sqlLiteInserts, consts.SqliteDatabase)
	if err != nil {
		fmt.Println("Error inserting drivers in database. Check the logs")
		logs.Errorf("Error insert in database. Error: %v", err)
	}

	// files created
	err = files.GenerateFile(sqlAcl, files.CreationFileRouteAclSql("ACL", "sql"))
	controlErrors(err)

	err = files.GenerateFile(jsonAcl, files.CreationFileRouteJson("usersCouchbase", "json"))
	controlErrors(err)
	err = files.GenerateFile(jsonEndPoint, files.CreationFileRouteAclJson("ACL-EP", "json"))
	controlErrors(err)

	err = files.GenerateFile(namesT, files.CreationFileRouteNames("names", "txt"))
	controlErrors(err)

	err = files.GenerateFile(sqlLiteInserts, files.CreationFileRouteSql("insertSQLIteQuery", "sql"))
	controlErrors(err)

	http.AuthEndpointCall(jsonEndPoint)

	fmt.Println("Finish")
}

func controlErrors(err error) {
	if err != nil {
		logs.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/lo")
	}
}
