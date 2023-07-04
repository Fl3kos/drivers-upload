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
	shop := shopCodes[0]

	phone := getPhones.GetAllPhones()[0]

	for {
		fmt.Printf("Are you secured to publish users with warehouse code %v and phone %v? (y/n) ", shop, phone)
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			break
		}

		if answer == "n" {
			fmt.Printf("The shop code or phone not the expected")
			log.Fatalf("The shop code or phone not the expected")
		}
	}

	users := generateUsers(list, shop, phone)
	finalJson := json.GenerateUsersJson(users)

	err := files.GenerateFile(finalJson, files.CreationFileUserList())

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	http.AuthEndpointCall(finalJson)

	var publish bool = false

	for {
		fmt.Printf("Are you publish roles to users with warehouse code %v? (y/n) ", shop)
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

	finalSql := generateSqlAndAclRole(users, shop, publish)
	err = files.GenerateFile(finalSql, files.CreationFileRouteAclSql("ACL", "sql"))
	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	fmt.Println("Finish")
}

func generateSqlAndAclRole(users []users.FUser, shopCode string, publish bool) string {
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
			log.Errorln("User dont identify")
		}

		if publish {
			//upload users to acl appPicking with role and store code
			token := strings.Split(files.ReadFile(files.ReadToken(consts.TokenFile)), "\n")[0]
			userAcl := json.GenerateAclJson(consts.AppPickingEnv, shopCode, rolePickingCode, false)
			go http.AclEndpointCall(userAcl, user.Username, token)

			//upload users to acl console with role and store code
			userAcl = json.GenerateAclJson(consts.ConsoleEnv, shopCode, roleConsoleCode, false)
			go http.AclEndpointCall(userAcl, user.Username, token)
		}
		usernames = append(usernames, user.Username)
	}

	finalSql = finalSql + "\n\n" + sql.GenerateAclInsert(usernames, appPickingRole)
	finalSql = finalSql + "\n" + sql.GenerateAclInsert(usernames, consoleRole)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, shopCode, consts.AppPickingEnv)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, shopCode, consts.ConsoleEnv)

	return finalSql
}

func generateUsers(list users.UsersList, warehouseCode, phoneNumber string) []users.FUser {
	var user users.FUser
	var users []users.FUser

	for i := 1; i <= list.Pkr; i++ {
		user = createUser(i, warehouseCode, "PKR", phoneNumber)
		users = append(users, user)
	}

	for i := 1; i <= list.Crd; i++ {
		user = createUser(i, warehouseCode, "CRD", phoneNumber)
		users = append(users, user)
	}

	for i := 1; i <= list.Adm; i++ {
		user = createUser(i, warehouseCode, "ADM", phoneNumber)
		users = append(users, user)
	}

	return users
}

func createUser(i int, warehouseCode string, userType string, phoneNumber string) users.FUser {
	var user users.FUser

	number := strconv.Itoa(i)

	for j := len(number); j < 3; j++ {
		number = "0" + number
	}

	for i := len(warehouseCode); i < 5; i++ {
		warehouseCode = "0" + warehouseCode
	}

	userAndPassword := fmt.Sprintf("%v%v%v", userType, warehouseCode, number)

	user.Username = userAndPassword
	user.Password = userAndPassword

	switch userType {
	case "PKR":
		user.Firstname = "APP PICKING"
		user.RoleCode = string(userType)
		break
	case "CRD":
		user.Firstname = "COORDINADOR"
		user.RoleCode = string(userType)
		break
	case "ADM":
		user.Firstname = "ADMINISTRADOR"
		user.RoleCode = string(userType)
		break
	default:
		log.Errorln("User dont identify")
	}

	lastname, err := numtoletter.IntLetra(i)
	if err != nil {

	}
	user.Lastname = strings.ToUpper(lastname)
	user.Email = fmt.Sprintf("t%ves@stores.diagroup.com", warehouseCode)
	user.Phone = phoneNumber

	return user
}
