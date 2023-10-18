package getPhones

import (
	"strings"
	files "support-utils/methods/file"
	"support-utils/structs/handlers"
)

func GetAllPhones() []string {
	phonesNumber := files.ReadFile(files.ReadFileRoute("phoneNumbers", "txt"))
	allPhones := strings.Split(phonesNumber, "\n")

	return allPhones
}

func DriversPhone(drivers []handlers.Driver)[]string {

	var phones []string

	for _, driver := range drivers {
		phone := driver.PhoneNumber

		phones = append(phones, phone)
	}

	return phones
}