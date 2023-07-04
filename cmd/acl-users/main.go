package main

import (
	"drivers-create/consts"
	files "drivers-create/methods/file"
	"drivers-create/methods/gets/getShops"
	"drivers-create/methods/http"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
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

	for {
		fmt.Printf("Are you secured to assign roles to users with warehouse code %v? (y/n) ", shop)
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

	pkr := generateUsers(list.Pkr, string(users.PKR), shop)
	crd := generateUsers(list.Crd, string(users.CRD), shop)
	adm := generateUsers(list.Adm, string(users.ADM), shop)

	errPkr := publishToAcl(pkr, shop, string(users.PKR))
	errCrd := publishToAcl(crd, shop, string(users.CRD))
	errAdm := publishToAcl(adm, shop, string(users.ADM))

	if errPkr != nil {
		log.Errorf("Error generating file: %v", errPkr)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	if errCrd != nil {
		log.Errorf("Error generating file: %v", errCrd)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	if errAdm != nil {
		log.Errorf("Error generating file: %v", errAdm)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	fmt.Println("Finish")
}

func publishToAcl(users []string, shopCode, role string) error {
	var appPickingRole string
	var consoleRole string
	appPickingEnv := "WMSPIC"
	consoleEnv := "ECOMUI"

	switch role {
	case "PKR":
		appPickingRole = "42"
		consoleRole = "46"
		break
	case "CRD":
		appPickingRole = "43"
		consoleRole = "47"
		break
	case "ADM":
		appPickingRole = "41"
		consoleRole = "45"
		break
	default:
		log.Errorln("User dont identify")
	}

	token := strings.Split(files.ReadFile(files.ReadToken(consts.TokenFile)), "\n")[0]

	//upload users to acl appPicking with role and store code
	userAcl := json.GenerateAclJson(appPickingEnv, shopCode, appPickingRole, false)
	for _, user := range users {
		http.AclEndpointCall(userAcl, user, token)
	}

	//upload users to acl console with role and store code
	userAcl = json.GenerateAclJson(consoleEnv, shopCode, consoleRole, false)
	for _, user := range users {
		http.AclEndpointCall(userAcl, user, token)
	}

	return nil
}

func generateUsers(cuantity int, userType, shopNumber string) []string {
	var users []string

	for i := 1; i <= cuantity; i++ {

		number := strconv.Itoa(i)

		for j := len(number); j < 3; j++ {
			number = "0" + number
		}

		for i := len(shopNumber); i < 5; i++ {
			shopNumber = "0" + shopNumber
		}
		user := fmt.Sprintf("%v%v%v", userType, shopNumber, number)
		users = append(users, user)
	}
	return users
}
