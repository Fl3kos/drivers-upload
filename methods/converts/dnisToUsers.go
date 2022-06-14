package converts

import (
	mt "drivers-create/methods"
	"strings"
)

func ConvertAllDnisToUsers(allDnis []string) []string {
	var allUsers = []string{}
	for _, dni := range allDnis {
		var user string
		var letter = dni[8:9]

		if mt.IsNumber(dni[0:1]) {
			user = letter + dni[1:8]
		} else {
			user = dni[:1] + letter + dni[2:8]
		}
		user = strings.ToUpper(user)
		allUsers = append(allUsers, user)
	}

	return allUsers
}
