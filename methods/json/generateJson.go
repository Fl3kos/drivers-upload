package json

import (
	"crypto/sha256"
	"drivers-create/consts"
	logs "drivers-create/methods/log"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateJson(allNames, allPasswords, allUsers []string) string {
	logs.Debugln("Generating Json")

	json := "[\n"

	for i, _ := range allPasswords {
		if allNames[i] != "" {
			firstname, lastname := getFirstNameAndLastName(allNames[i])
			encodedPassword := encodePassword(allPasswords[i])

			value := generateJson(allUsers[i], encodedPassword, firstname, lastname)

			if i != len(allPasswords)-1 {
				value = value + ",\n"
			}

			json = json + value
		}
	}

	json = json + "\n]"

	return json
}

func generateJson(username, password, firstname, lastname string) string {
	logs.Debugln("Generating JSON to", username)

	json :=
		`	{
		"username" : "%v",
		"password" : "%v",
		"firstname" : "%v",
		"lastname" : "%v",
		"collection" : "%v",
		"userType": "%v"
	}`

	json = fmt.Sprintf(json, username, password, firstname, lastname, consts.Collection_Json, consts.UserType_Json)

	logs.Debugln("JSON generated")

	return json
}

func GenerateAclJson(allNames, allPasswords, allUsers, allPhones []string) string {
	logs.Debugln("Generating ACL Json")

	json := "[\n"

	for i, _ := range allPasswords {
		if allNames[i] != "" {
			firstname, lastname := getFirstNameAndLastName(allNames[i])
			encodedPassword := encodePassword(allPasswords[i])

			value := generateAclJson(allUsers[i], encodedPassword, firstname, lastname, allPhones[i])

			if i != len(allPasswords)-1 {
				value = value + ",\n"
			}

			json = json + value
		}
	}

	json = json + "\n]"

	return json
}

func GenerateAclJsonDontEncript(allNames, allPasswords, allUsers, allPhones []string) string {
	logs.Debugln("Generating ACL Json")

	json := "[\n"

	for i, _ := range allPasswords {
		if allNames[i] != "" {
			firstname, lastname := getFirstNameAndLastName(allNames[i])

			value := generateAclJson(allUsers[i], allPasswords[i], firstname, lastname, allPhones[i])

			if i != len(allPasswords)-1 {
				value = value + ",\n"
			}

			json = json + value
		}
	}

	json = json + "\n]"

	return json
}

func generateAclJson(username, password, firstname, lastname, phone string) string {
	logs.Debugln("Generating ACL JSON to", username)

	json :=
		`	{
		"email": "%v",
		"firstname": "%v",
		"lastname": "%v",
		"password": "%v",
		"phone": "%v",
		"username": "%v"
	}`

	json = fmt.Sprintf(json, consts.GenericDriverEmail, firstname, lastname, password, phone, username)

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
