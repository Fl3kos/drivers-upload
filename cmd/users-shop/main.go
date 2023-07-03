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

	pkr := generateUsers(list.Pkr, string(users.PKR), shop, phone)
	crd := generateUsers(list.Crd, string(users.CRD), shop, phone)
	adm := generateUsers(list.Adm, string(users.ADM), shop, phone)

	finalJson := json.GenerateUsersJson(pkr, crd, adm)

	err := files.GenerateFile(finalJson, files.CreationFileUserList())

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	http.AuthEndpointCall(finalJson)

	sqlPkr := generateSql(pkr, shop, string(users.PKR))
	sqlCrd := generateSql(crd, shop, string(users.CRD))
	sqlAdm := generateSql(adm, shop, string(users.ADM))
	finalSql := sqlPkr + sqlCrd + sqlAdm

	err = files.GenerateFile(finalSql, files.CreationFileRouteAclSql("ACL", "sql"))
	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	fmt.Println("Finish")
}

func generateSql(users []users.User, shopCode, role string) string {
	var usernames []string
	var appPickingRole string
	var consoleRole string
	var rolePickingCode string
	var roleConsoleCode string
	appPickingEnv := "WMSPIC"
	consoleEnv := "ECOMUI"

	finalSql := ""

	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	switch role {
	case "PKR":
		appPickingRole = "ROLE_WMSPIC_PICKER"
		consoleRole = "ROLE_ECOMUI_WMS_PICKER"
		rolePickingCode = "42"
		roleConsoleCode = "46"
		break
	case "CRD":
		appPickingRole = "ROLE_WMSPIC_COORDINATOR"
		consoleRole = "ROLE_ECOMUI_WMS_COORDINATOR"
		rolePickingCode = "43"
		roleConsoleCode = "47"
		break
	case "ADM":
		appPickingRole = "ROLE_WMSPIC_ADMIN"
		consoleRole = "ROLE_ECOMUI_WMS_ADMIN"
		rolePickingCode = "41"
		roleConsoleCode = "45"
		break
	default:
		log.Errorln("User dont identify")
	}

	var publish bool = false

	for {
		fmt.Printf("Are you publish roles to users with warehouse code %v? (y/n) ", shopCode)
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

	if publish {
		//upload users to acl appPicking with role and store code
		token := strings.Split(files.ReadFile(files.ReadToken(consts.TokenFile)), "\n")[0]
		userAcl := json.GenerateAclJson(appPickingEnv, shopCode, rolePickingCode)
		for _, user := range usernames {
			http.AclEndpointCall(userAcl, user, token)
		}

		//upload users to acl console with role and store code
		userAcl = json.GenerateAclJson(consoleEnv, shopCode, roleConsoleCode)
		for _, user := range usernames {
			http.AclEndpointCall(userAcl, user, token)
		}
	}
	finalSql = finalSql + "\n\n" + sql.GenerateAclInsert(usernames, appPickingRole)
	finalSql = finalSql + "\n" + sql.GenerateAclInsert(usernames, consoleRole)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, shopCode, appPickingEnv)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, shopCode, consoleEnv)

	return finalSql
}

func generateUsers(cuantity int, userType, shopNumber, phoneNumber string) []users.User {
	var user users.User
	var users []users.User

	for i := 1; i <= cuantity; i++ {

		number := strconv.Itoa(i)

		for j := len(number); j < 3; j++ {
			number = "0" + number
		}

		for i := len(shopNumber); i < 5; i++ {
			shopNumber = "0" + shopNumber
		}

		userAndPassword := fmt.Sprintf("%v%v%v", userType, shopNumber, number)

		user.Username = userAndPassword
		user.Password = userAndPassword

		switch userType {
		case "PKR":
			user.Firstname = "APP PICKING"
			break
		case "CRD":
			user.Firstname = "COORDINADOR"
			break
		case "ADM":
			user.Firstname = "ADMINISTRADOR"
			break
		default:
			log.Errorln("User dont identify")
		}

		lastname, err := numtoletter.IntLetra(i)
		if err != nil {

		}
		user.Lastname = strings.ToUpper(lastname)
		user.Email = fmt.Sprintf("t%ves@stores.diagroup.com", shopNumber)
		user.Phone = phoneNumber

		users = append(users, user)
	}
	return users
}
