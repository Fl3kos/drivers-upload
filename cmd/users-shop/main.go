package main

import (
	files "drivers-create/methods/file"
	"drivers-create/methods/gets/getPhones"
	"drivers-create/methods/gets/getShops"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	numtoletter "drivers-create/methods/numToLetter"
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
	fmt.Println(finalJson)
	err := files.GenerateFile(finalJson, files.CreationFileUserList())

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/lo")
	}

	fmt.Println("Finish")
}

func generateUsers(cuantity int, userType, shopNumber, phoneNumber string) []users.User {
	var user users.User
	var users []users.User

	for i := 1; i <= cuantity; i++ {

		number := strconv.Itoa(i)

		//PKR-00937-002
		for j := len(number); j < 3; j++ {
			number = "0" + number
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
