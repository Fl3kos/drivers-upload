package dniToUser

import (
	"strings"
	mt "support-utils/methods"
	"support-utils/structs/handlers"
)

func ConvertAllDnisToUsers(allDnis []string) []string {
	var allUsers = []string{}

	for _, dni := range allDnis {
		user := ""
		if dni != "" {

			var letter = dni[8:9]

			if mt.IsNumber(dni[0:1]) {
				user = letter + dni[1:8]
			} else {
				user = dni[:1] + letter + dni[2:8]
			}
			user = strings.ToUpper(user)

		}
		allUsers = append(allUsers, user)

	}

	return allUsers
}

func ConvertDnisToUsernames(drivers []handlers.Driver) []string {
	var usernames = []string{}

	for _, driver := range drivers {
		user := ""
		dni := driver.Dni

		var letter = dni[8:9]

		if mt.IsNumber(dni[0:1]) {
			user = letter + dni[1:8]
		} else {
			user = dni[:1] + letter + dni[2:8]
		}

		user = strings.ToUpper(user)
		usernames = append(usernames, user)
		}



	return usernames
}