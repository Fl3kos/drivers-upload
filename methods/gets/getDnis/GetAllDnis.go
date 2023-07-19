package getDnis

import (
	"strings"
	files "support-utils/methods/file"
)

func GetAllDnis() []string {
	dnis := files.ReadFile(files.ReadFileRoute("dnis", "txt"))
	dnis = strings.ToUpper(dnis)

	substring := dnis[:len(dnis)-1]
	allDnis := strings.Split(substring, "\n")

	//make trim to all dnis
	for _, dni := range allDnis {
		if dni != "" {
			dni = strings.TrimSpace(dni)
		}

	}

	return allDnis
}
