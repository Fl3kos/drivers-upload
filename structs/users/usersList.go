package users

import (
	"encoding/json"
	"support-utils/methods/file"
)

type UsersList struct {
	Pkr int
	Crd int
	Adm int
}

func ReturnList() UsersList {
	var list UsersList
	jsonS := file.ReadFile(file.ReadUserListFile())
	json.Unmarshal([]byte(jsonS), &list)

	return list
}
