package converts

import (
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

	return allNamesT
}

func UsersAndPasswords(allNames, allUsers, allPasswords []string) string {
	var usersAndPasswords = "NAME\n"

	for _, name := range allNames {
		usersAndPasswords = usersAndPasswords + name + "\n"
	}

	usersAndPasswords = usersAndPasswords + "\nUSERS\n"

	for _, user := range allUsers {
		usersAndPasswords = usersAndPasswords + user + "\n"
	}

	usersAndPasswords = usersAndPasswords + "\nPASSWORDS\n"
	for _, password := range allPasswords {
		usersAndPasswords = usersAndPasswords + password + "\n"
	}

	return usersAndPasswords
}
