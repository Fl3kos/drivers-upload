package json

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	logs "drivers-create/methods/log"
)

func GenerateJson(allNames, allPasswords, allUsers []string) string {
	logs.DebugLog.Println("Generating Json")

	json := "[\n"

	m := len(allPasswords)

	collection := "\"authentication\""
	userType := "\"userType\": \"ECOMMERCE_USER\""
	for i := 0; i < m; i++ {
		name := strings.Split(allNames[i], " ")
		firstname := "\"" + name[0] + "\""
		lastname := "\""

		for j := 1; j < len(name); j++ {
			lastname = lastname + name[j] + " "
		}

		lastname = strings.TrimSpace(lastname)
		lastname = lastname + "\""

		ps := encodePassword(allPasswords[i])

		user, pass := formatUserAndPassword(allUsers[i], ps)

		value := "\t{\n\t\t\"username\" : " + user + " ,\n\t\t\"password\" : " + pass + " ,\n\t\t\"firstname\" : " + firstname + " ,\n\t\t\"lastname\" : " + lastname + " ,\n\t\t\"collection\" : " + collection + ",\n\t\t" + userType + "\n\t}"

		if i != m-1 {
			value = value + ",\n"
		}

		json = json + value
	}

	json = json + "\n]"

	return json
}

//encode the password
func encodePassword(password string) string {
	logs.DebugLog.Println("Password is encripting")
	encrypted := sha256.Sum256([]byte(password))
	ps := hex.EncodeToString(encrypted[:])

	logs.DebugLog.Println("Password was encripted succesfull")

	return ps
}

//this method format the user and password for the json text
func formatUserAndPassword(user, password string) (userF string, passF string) {
	userF = "\"" + user + "\""
	passF = "\"" + password + "\""
	return
}
