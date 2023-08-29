package createUsers

import (
	"fmt"
	"support-utils/methods"
	dniM "support-utils/methods/dni"
	"support-utils/methods/dniToUser"
	files "support-utils/methods/file"
	"support-utils/methods/gets/getNames"
	"support-utils/methods/gets/getPhones"
	//"support-utils/methods/http"
	"support-utils/methods/json"
	logs "support-utils/methods/log"
	"support-utils/methods/userToPassword"
	"support-utils/structs/handlers"
	"support-utils/structs/responses"
)

func CreateDrivers (drivers []handlers.Driver, warehouse string) ([]responses.DriverResponse, []string) {
	//TODO:


	dnisIncorrect, err := dniM.ComprobeDriversDnis(drivers)
	if err != nil {
		fmt.Println("Error without validate dnis, check the logs")
		logs.Errorf("Error validating dnis, incorrect DNIs: %v. Error: %v", dnisIncorrect, err)
		return nil, dnisIncorrect
	}
	//convert dnis to username
	//convert usernames to password


	usernames := dniToUser.ConvertDnisToUsernames(drivers)
	passwords := userToPassword.ConvertAllUsersToPasswords(usernames)

	names := getNames.DriversName(drivers)
	phones := getPhones.DriversPhone(drivers)


	var response []responses.DriverResponse

	for i, _ := range drivers {
		driverResponse := responses.DriverResponse{Name: names[i],Dni: drivers[i].Dni, Username: usernames[i],Phone: phones[i], Password: passwords[i]}

		response = append(response, driverResponse)
	}



	jsonEndPoint := json.AuthEndpointJson(names, passwords, usernames, phones, warehouse)

	err = files.GenerateFile(jsonEndPoint, files.CreationFileRouteAclJson("ACL-EP", "json"))
	methods.ControlErrors(err)

	//http.AuthEndpointCall(jsonEndPoint)
	//TODO create nektria driver without endpoint

	return response, dnisIncorrect
}

