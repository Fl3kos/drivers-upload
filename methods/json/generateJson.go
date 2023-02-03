package json

import (
	"crypto/sha256"
	logs "drivers-create/methods/log"
	"drivers-create/structs/users"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

func GenerateJson(allNames, allPasswords, allUsers, allPhones, allShops []string) string {
	logs.Debugln("Generating ACL Json")

	json := "[\n"
	s := 0
	shop := allShops[s]

	for i, _ := range allPasswords {
		if allNames[i] != "" {
			firstname, lastname := getFirstNameAndLastName(allNames[i])
			encodedPassword := encodePassword(allPasswords[i])
			email := fmt.Sprintf("t%ves@stores.diagroup.com", shop)

			value := generateJson(allUsers[i], encodedPassword, firstname, lastname, allPhones[i], email)

			if i != len(allPasswords)-1 {
				value = value + ",\n"
			}

			json = json + value
		} else {
			s = s + 1
			shop = allShops[s]
		}
	}

	json = json + "\n]"

	return json
}

func GenerateUsersJson(pkr, crd, adm []users.User) string {
	finalJson := "{\n\t\"users\" : ["
	pkrJson := generateUsersJson(pkr)
	crdJson := generateUsersJson(crd)
	admJson := generateUsersJson(adm)

	finalJson = finalJson + pkrJson + crdJson + admJson

	finalJson = finalJson[:len(finalJson)-1]
	finalJson = finalJson + "\n\t]\n}"
	return finalJson
}

func generateUsersJson(users []users.User) string {
	finalJson := ""

	for _, user := range users {
		jsonByte, _ := json.Marshal(user)
		jsonTxt := string(jsonByte)
		finalJson = finalJson + "\n\t\t" + jsonTxt + ","
	}

	return finalJson
}

func GenerateEndpointJson(allNames, allPasswords, allUsers, allPhones, allShops []string) string {
	logs.Debugln("Generating ACL Json")

	json := "{\n\t\"user\": [\n"
	s := 0
	shop := allShops[s]

	for i, _ := range allPasswords {
		if allNames[i] != "" {
			firstname, lastname := getFirstNameAndLastName(allNames[i])
			email := fmt.Sprintf("t%ves@stores.diagroup.com", shop)
			value := generateEndpointJson(allUsers[i], allPasswords[i], firstname, lastname, allPhones[i], email)

			if i != len(allPasswords)-1 {
				value = value + ",\n"
			}

			json = json + value
		} else {
			s = s + 1
			shop = allShops[s]
		}
	}

	json = json + "\n\t]\n}"

	return json
}

func generateEndpointJson(username, password, firstname, lastname, phone, email string) string {
	logs.Debugln("Generating ACL JSON to", username)

	json :=
		`		{
			"email": "%v",
			"firstname": "%v",
			"lastname": "%v",
			"password": "%v",
			"phone": "%v",
			"username": "%v"
		}`

	json = fmt.Sprintf(json, email, firstname, lastname, password, phone, username)

	logs.Debugln("ACL JSON generated")

	return json
}

func generateJson(username, password, firstname, lastname, phone, email string) string {
	logs.Debugln("Generating ACL JSON to", username)

	json :=
		`	{
		"email": "%v",
		"firstname": "%v",
		"lastname": "%v",
		"collection": "authentication",
		"password": "%v",
		"phone": "%v",
		"username": "%v",
		"userType": "ECOMMERCE_USER"
	}`

	json = fmt.Sprintf(json, email, firstname, lastname, password, phone, username)

	logs.Debugln("ACL JSON generated")

	return json
}

func getFirstNameAndLastName(completeName string) (string, string) {
	name := strings.Split(completeName, " ")
	firstname := name[0]

	lastname := ""
	for j := 1; j < len(name); j++ {
		lastname = lastname + name[j] + " "
	}
	lastname = strings.TrimSpace(lastname)

	return firstname, lastname
}

func encodePassword(password string) string {
	logs.Debugln("Password is encripting")

	encrypted := sha256.Sum256([]byte(password))
	encodedPassword := hex.EncodeToString(encrypted[:])

	logs.Debugln("Password was encripted succesfull")

	return encodedPassword
}
