package converts

import (
	mt "drivers-create/methods"
	"fmt"
	"strings"
)

func ConvertAllDnisToUsers(allDnis []string) []string {
	var allUsers = []string{}
	fmt.Println(len(allDnis))
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
