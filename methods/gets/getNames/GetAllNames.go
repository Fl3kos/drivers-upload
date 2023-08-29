package getNames

import (
	"strings"
	files "support-utils/methods/file"
	"support-utils/structs/handlers"
)

func GetAllNames() []string {
	names := files.ReadFile(files.ReadFileRoute("names", "txt"))
	allNames := strings.Split(names, "\n")

	//make trim to all users
	for i, name := range allNames {
		if name != "" {
			allNames[i] = strings.TrimSpace(name)
		}
	}

	return allNames
}

func DriversName(drivers []handlers.Driver)[]string {

	var names []string

	for _, driver := range drivers {
		name := driver.Name

		names = append(names, name)
	}

	return names
}