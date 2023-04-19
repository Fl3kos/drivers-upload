package main

import (
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

	pkr := generateUsers(list.Pkr, string(users.PKR), shop, phone)
	crd := generateUsers(list.Crd, string(users.CRD), shop, phone)
	adm := generateUsers(list.Adm, string(users.ADM), shop, phone)

	finalJson := json.GenerateUsersJson(pkr, crd, adm)
	sqlPkr := generateSql(pkr, shop, string(users.PKR))
	sqlCrd := generateSql(crd, shop, string(users.CRD))
	sqlAdm := generateSql(adm, shop, string(users.ADM))
	finalSql := sqlPkr + sqlCrd + sqlAdm

	err := files.GenerateFile(finalJson, files.CreationFileUserList())

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	err = files.GenerateFile(finalSql, files.CreationFileRouteAclSql("ACL", "sql"))

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	http.AuthEndpointCall(finalJson)

	fmt.Println("Finish")
}

func generateSql(users []users.User, shopCode, role string) string {
	var usernames []string
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}
	finalSql := ""
	role1 := ""
	role2 := ""
	environoment1 := "WMSPIC"
	environoment2 := "ECOMUI"
	switch role {
	case "PKR":
		role1 = "ROLE_WMSPIC_PICKER"
		role2 = "ROLE_ECOMUI_WMS_PICKER"
		break
	case "CRD":
		role1 = "ROLE_WMSPIC_COORDINATOR"
		role2 = "ROLE_ECOMUI_WMS_COORDINATOR"
		break
	case "ADM":
		role1 = "ROLE_WMSPIC_ADMIN"
		role2 = "ROLE_ECOMUI_WMS_ADMIN"
		break
	default:
		log.Errorln("User dont identify")
	}
	finalSql = finalSql + "\n\n" + sql.GenerateAclInsert(usernames, role1)
	finalSql = finalSql + "\n" + sql.GenerateAclInsert(usernames, role2)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, shopCode, environoment1)
	finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, shopCode, environoment2)

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
