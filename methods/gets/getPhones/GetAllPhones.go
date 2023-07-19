package getPhones

import (
	"strings"
	files "support-utils/methods/file"
)

func GetAllPhones() []string {
	phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
	allPhones := strings.Split(phonesNumber, "\n")

	return allPhones
}
