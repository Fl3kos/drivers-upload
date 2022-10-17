package getPhones

import (
	files "drivers-create/methods/file"
	"strings"
)

func GetAllPhones() []string {
	phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
	allPhones := strings.Split(phonesNumber, "\n")

	return allPhones
}
