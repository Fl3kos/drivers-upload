package getNames

import (
	files "drivers-create/methods/file"
	"strings"
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

