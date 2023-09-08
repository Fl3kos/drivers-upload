package createUsers

import (
	"support-utils/methods/acl"
	"support-utils/methods/http"
	"support-utils/methods/json"
	"support-utils/structs/handlers"
)

func CreateWarehouseUsers(users handlers.WarehouseUsers, warehouse, authorization string) error {
	aclUsers := acl.GenerateUsersApi(users, warehouse)
	finalJson := json.GenerateUsersJson(aclUsers)

	http.AuthEndpointCall(finalJson)

	acl.PublishAclUsersApi(aclUsers, warehouse, authorization)
	return nil
}
