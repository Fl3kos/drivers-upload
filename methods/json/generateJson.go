package json

import (
	"crypto/sha256"
	logs "drivers-create/methods/log"
	"drivers-create/structs/users"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
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

func GenerateUsersJson(aclUsers []users.AclUser) string {
	finalJson := "{\n\t\"user\" : ["
	json := generateUsersJson(aclUsers)

	finalJson = finalJson + json
	finalJson = finalJson[:len(finalJson)-1]
	finalJson = finalJson + "\n\t]\n}"

	return finalJson
}

func generateUsersJson(aclUsers []users.AclUser) string {
	finalJson := ""

	for _, fuser := range aclUsers {
		user := users.UserConstruct(fuser.Email, fuser.Firstname, fuser.Lastname, fuser.Password, fuser.Phone, fuser.Username)

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

func GenerateSorterMap(locationAndSorter []*xlsx.Row, warehouseCode string) string {
	var sorter []string
	var locations []string

	for i, ubication := range locationAndSorter {
		if i > 0 {
			if ubication.Cells[9].String() == "" {
				return ""
			}

			ubi := strings.Split(ubication.Cells[9].String(), ".")[1]

			if ubi[:1] == "0" {
				fmt.Println(ubi[:1])
				ubi = ubi[1:2]
			}

			sorter = append(sorter, ubi)
			locations = append(locations, ubication.Cells[11].String())
		}
	}

	sorterMap := generateSorterMap(sorter, locations, warehouseCode)
	return sorterMap
}

func generateSorterMap(sorter, location []string, warehouse string) string {
	var sorterMap string

	sorterMap =
		`{
		"store_code" : "%v",
		"yard_chute_map": {
			%v
		}
	}`
	values := ""
	value := "\"%v\":\"%v\""

	for i, _ := range sorter {
		valueF := fmt.Sprintf(value, location[i], sorter[i])

		if i != len(sorter)-1 {
			valueF = valueF + ",\n"
		}
		values = values + valueF
	}

	sorterMap = fmt.Sprintf(sorterMap, warehouse, values)

	return sorterMap
}

func GenerateAclJson(aplicationCode, storeCode, roleCode string, isDriver bool) string {
	logs.Debugln("Generating ACL Json")
	var segmentation string = ""

	if !isDriver {
		segmentation = `{
		"type": "AND",
		"content": [
		  {
			"dimension": "storeCode",
			"values": [
			  "%v"
			]
		  },
		  {
			"dimension": "country",
			"values": [
			  "ES"
			]
		  }
		]
	  }`

		segmentation = fmt.Sprintf(segmentation, storeCode)
	}

	json :=
		`{
			"userType": "ECOMMERCE_USER",
			"rolesForApplications": [
			  {
				"applicationCode": "%v",
				"roleCodes": [
				  %v
				]
			  }
			],
			"segmentationsForApplications": [
			  {
				"applicationCode": "%v",
				"segmentation": [%v]
			  }
			]
		  }`

	json = fmt.Sprintf(json, aplicationCode, roleCode, aplicationCode, segmentation)

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
