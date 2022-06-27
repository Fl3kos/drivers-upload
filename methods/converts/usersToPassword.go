package converts

import (
	common "drivers-create/methods"
	"strings"
)

func ConvertAllUsersToPasswords(allUsers []string) []string {
	var allPasswords = []string{}
	var letter string

	for _, user := range allUsers {
		password := ""
		if user != "" {
			if common.IsNumber(user[1:2]) {
				letter = strings.ToLower(user[0:1])
			} else {
				letter = strings.ToLower(user[1:2])
			}

			password = user[:7] + letter

		}
		allPasswords = append(allPasswords, password)
	}
	return allPasswords
}
