package main

import (
	"drivers-create/consts"
	files "drivers-create/methods/file"
	"drivers-create/methods/gets/getPhones"
	"drivers-create/methods/gets/getShops"
	"drivers-create/methods/http"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	numtoletter "drivers-create/methods/numToLetter"
	"drivers-create/methods/sql"
	"drivers-create/structs/users"

	"strconv"

	"fmt"
	"strings"
)

func main() {
	log.InitLogger()
	list := users.ReturnList()
	shops := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
	shops = shops[:len(shops)-1]

	shopCodes, _ := getShops.GetShopCodesAndShopNames(shops)
	warehouseCode := shopCodes[0]

	phone := getPhones.GetAllPhones()[0]

	for {
		fmt.Printf("Are you secured to publish aclUsers with warehouse code %v and phone %v? (y/n) ", warehouseCode, phone)
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			break
		}

		if answer == "n" {
			fmt.Printf("The warehouseCode code or phone not the expected")
			log.Fatalf("The warehouseCode code or phone not the expected")
		}
	}

	aclUsers := generateUsers(list, warehouseCode, phone)
	finalJson := json.GenerateUsersJson(aclUsers)

	err := files.GenerateFile(finalJson, files.CreationFileUserList())

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	http.AuthEndpointCall(finalJson)

	var publish bool = false

	for {
		fmt.Printf("Are you publish roles to aclUsers with warehouse code %v? (y/n) ", warehouseCode)
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			publish = true
		}

		if answer == "n" {
			log.Debugln("Roles not publish")
			publish = false
		}

		break
	}

	finalSql := generateSqlAndAclRole(aclUsers, warehouseCode, publish)
	err = files.GenerateFile(finalSql, files.CreationFileRouteAclSql("ACL", "sql"))
	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	fmt.Println("Finish")
}

func generateSqlAndAclRole(users []users.AclUser, warehouseCode string, publish bool) string {
	var usernames []string
	var appPickingRole string
	var consoleRole string
	var rolePickingCode string
	var roleConsoleCode string

	finalSql := ""

	for _, user := range users {

		switch user.RoleCode {
		case "PKR":
			appPickingRole = consts.AppPickingRolePkr
			consoleRole = consts.ConsoleRolePkr
			rolePickingCode = consts.PickingCodePkr
			roleConsoleCode = consts.ConsoleCodePkr
			break
		case "CRD":
			appPickingRole = consts.AppPickingRoleCrd
			consoleRole = consts.ConsoleRoleCrd
			rolePickingCode = consts.PickingCodeCrd
			roleConsoleCode = consts.ConsoleCodeCrd
			break
		case "ADM":
			appPickingRole = consts.AppPickingRoleAdm
			consoleRole = consts.ConsoleRoleAdm
			rolePickingCode = consts.PickingCodeAdm
			roleConsoleCode = consts.ConsoleCodeAdm
			break
		default:
			log.Errorln("AuthUser dont identify")
		}

		if publish {
			//upload users to acl appPicking with role and store code
			token := strings.Split(files.ReadFile(files.ReadToken(consts.TokenFile)), "\n")[0]
			userAcl := json.GenerateAclJson(consts.AppPickingEnv, warehouseCode, rolePickingCode, false)
			go http.AclEndpointCall(userAcl, user.Username, token)

			//upload users to acl console with role and store code
			userAcl = json.GenerateAclJson(consts.ConsoleEnv, warehouseCode, roleConsoleCode, false)
			go http.AclEndpointCall(userAcl, user.Username, token)
		}
		usernames = append(usernames, user.Username)
	}

	finalSql = finalSql + "\n\n" + sql.GenerateAclInsert(usernames, appPickingRole)
	finalSql = finalSql + "\n" + sql.GenerateAclInsert(usernames, consoleRole)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, warehouseCode, consts.AppPickingEnv)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, warehouseCode, consts.ConsoleEnv)

	return finalSql
}

func generateUsers(list users.UsersList, warehouseCode, phoneNumber string) []users.AclUser {
	var aclUser users.AclUser
	var aclUsers []users.AclUser

	for i := 1; i <= list.Pkr; i++ {
		aclUser = createUser(i, warehouseCode, "PKR", phoneNumber)
		aclUsers = append(aclUsers, aclUser)
	}

	for i := 1; i <= list.Crd; i++ {
		aclUser = createUser(i, warehouseCode, "CRD", phoneNumber)
		aclUsers = append(aclUsers, aclUser)
	}

	for i := 1; i <= list.Adm; i++ {
		aclUser = createUser(i, warehouseCode, "ADM", phoneNumber)
		aclUsers = append(aclUsers, aclUser)
	}

	return aclUsers
}

func createUser(i int, warehouseCode string, userType string, phoneNumber string) users.AclUser {
	var aclUser users.AclUser

	number := strconv.Itoa(i)

	for j := len(number); j < 3; j++ {
		number = "0" + number
	}

	for i := len(warehouseCode); i < 5; i++ {
		warehouseCode = "0" + warehouseCode
	}

	userAndPassword := fmt.Sprintf("%v%v%v", userType, warehouseCode, number)

	aclUser.Username = userAndPassword
	aclUser.Password = userAndPassword

	switch userType {
	case "PKR":
		aclUser.Firstname = "APP PICKING"
		aclUser.RoleCode = string(userType)
		break
	case "CRD":
		aclUser.Firstname = "COORDINADOR"
		aclUser.RoleCode = string(userType)
		break
	case "ADM":
		aclUser.Firstname = "ADMINISTRADOR"
		aclUser.RoleCode = string(userType)
		break
	default:
		log.Errorln("AuthUser dont identify")
	}

	lastname, err := numtoletter.IntLetra(i)
	if err != nil {

	}
	aclUser.Lastname = strings.ToUpper(lastname)
	aclUser.Email = fmt.Sprintf("t%ves@stores.diagroup.com", warehouseCode)
	aclUser.Phone = phoneNumber

	return aclUser
}
