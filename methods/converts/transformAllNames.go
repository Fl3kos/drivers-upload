package converts

import (
	"fmt"
	"strings"
)

func TransformAllNames(allNames []string) string {
	var allNamesT string
	for _, name := range allNames {
		var nameF = strings.ToLower(name)
		//remplazar espacios por guiones
		nameF = strings.ReplaceAll(nameF, " ", "-")
		allNamesT = allNamesT + nameF + "\n"
	}

	fmt.Println("allNamesT:", allNamesT)
	return allNamesT
}

func UsersAndPasswords(allNames, allUsers, allPasswords []string, m int) string {
	var usersAndPasswords = "NAME\n"
	for i := 0; i < m; i++ {
		usersAndPasswords = usersAndPasswords + allNames[i] + "\n"
	}

	usersAndPasswords = usersAndPasswords + "\nUSERS\n"
	for i := 0; i < m; i++ {
		usersAndPasswords = usersAndPasswords + allUsers[i] + "\n"
	}

	usersAndPasswords = usersAndPasswords + "\nPASSWORDS\n"
	for i := 0; i < m; i++ {
		usersAndPasswords = usersAndPasswords + allPasswords[i] + "\n"
	}

	return usersAndPasswords
}
