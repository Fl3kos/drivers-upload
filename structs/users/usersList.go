package users

import (
	"drivers-create/methods/file"
	"encoding/json"
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
