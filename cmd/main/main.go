package main

import (
	"drivers-create/consts"
	"drivers-create/methods"
	"drivers-create/methods/acl"
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

	warehouses := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
	warehouses = warehouses[:len(warehouses)-1]

	shopCodes, shopNames := getShops.GetShopCodesAndShopNames(warehouses)

	csv.ExportDriversToCsv(allUsers, allNames, allPasswords, shopNames)

	// Create Json files
	jsonAcl := json.GenerateJson(allNames, allPasswords, allUsers, allPhones, shopCodes)
	jsonEndPoint := json.GenerateEndpointJson(allNames, allPasswords, allUsers, allPhones, shopCodes)

	// Create Names file
	namesT := convert.TransformAllNames(allNames)

	// Create SQL files
	driversInsert := sql.GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones)
	relationsInsert := sql.GenerateSqlLiteInsertRelationTable(allDnis, shopCodes)
	sqlAcl := sql.GenerateAclInsert(allUsers, consts.DriverRole)
	sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

	//insert in sqlite
	err = sqlite.InsertSqlite(sqlLiteInserts, consts.SqliteDatabase)
	if err != nil {
		fmt.Println("Error inserting drivers in database. Check the logs")
		logs.Errorf("Error insert in database. Error: %v", err)
	}

	// files created
	err = files.GenerateFile(sqlAcl, files.CreationFileRouteAclSql("ACL", "sql"))
	methods.ControlErrors(err)

	err = files.GenerateFile(jsonAcl, files.CreationFileRouteJson("usersCouchbase", "json"))
	methods.ControlErrors(err)

	err = files.GenerateFile(jsonEndPoint, files.CreationFileRouteAclJson("ACL-EP", "json"))
	methods.ControlErrors(err)

	err = files.GenerateFile(namesT, files.CreationFileRouteNames("names", "txt"))
	methods.ControlErrors(err)

	err = files.GenerateFile(sqlLiteInserts, files.CreationFileRouteSql("insertSQLIteQuery", "sql"))
	methods.ControlErrors(err)

	http.AuthEndpointCall(jsonEndPoint)

	for {
		fmt.Printf("Are you publish roles to drivers? (y/n)")
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			acl.PublisDrivershRoles(allUsers)
			break
		}

		if answer == "n" {
			logs.Debugln("Roles not publish")
			break
		}


	}

	//TODO Create selenium to publish nektria users

	fmt.Println("Finish")
}


