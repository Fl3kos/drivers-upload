package main

import (
	"fmt"
	"strings"
	"support-utils/consts"
	"support-utils/methods"
	"support-utils/methods/acl"
	convert "support-utils/methods/converts"
	csv "support-utils/methods/csv"
	dniM "support-utils/methods/dni"
	"support-utils/methods/dniToUser"
	files "support-utils/methods/file"
	"support-utils/methods/gets/getDnis"
	"support-utils/methods/gets/getNames"
	"support-utils/methods/gets/getPhones"
	"support-utils/methods/gets/getShops"
	"support-utils/methods/http"
	json "support-utils/methods/json"
	logs "support-utils/methods/log"
	sql "support-utils/methods/sql"
	"support-utils/methods/userToPassword"
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
	sqlAcl := sql.GenerateAclInsert(allUsers, consts.DriverRole)

	//insert in sqlite

	// files created
	err = files.GenerateFile(sqlAcl, files.CreationFileRouteAclSql("ACL", "sql"))
	methods.ControlErrors(err)

	err = files.GenerateFile(jsonAcl, files.CreationFileRouteJson("usersCouchbase", "json"))
	methods.ControlErrors(err)

	err = files.GenerateFile(jsonEndPoint, files.CreationFileRouteAclJson("ACL-EP", "json"))
	methods.ControlErrors(err)

	err = files.GenerateFile(namesT, files.CreationFileRouteNames("names", "txt"))
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
