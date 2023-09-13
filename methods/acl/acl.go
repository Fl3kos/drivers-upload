package acl

import (
	"strings"
	"support-utils/consts"
	files "support-utils/methods/file"
	"support-utils/methods/http"
	"support-utils/methods/json"
	"support-utils/methods/log"
	"support-utils/methods/sql"
	"support-utils/structs/handlers"
	"support-utils/structs/users"
)

func GenerateSql(users []users.AclUser, warehouseCode string) string {
	var appPickingRole string
	var consoleRole string

	finalSql := ""

	for _, user := range users {
		var usernames []string

		switch user.RoleCode {
		case "PKR":
			appPickingRole = consts.AppPickingRolePkr
			consoleRole = consts.ConsoleRolePkr
			break
		case "CRD":
			appPickingRole = consts.AppPickingRoleCrd
			consoleRole = consts.ConsoleRoleCrd
			break
		case "ADM":
			appPickingRole = consts.AppPickingRoleAdm
			consoleRole = consts.ConsoleRoleAdm
			break
		default:
			log.Errorln("AuthUser dont identify")
		}

		usernames = append(usernames, user.Username)

		finalSql = finalSql + "\n\n" + sql.GenerateAclInsert(usernames, appPickingRole)
		finalSql = finalSql + "\n" + sql.GenerateAclInsert(usernames, consoleRole)
		finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, warehouseCode, consts.AppPickingEnv)
		finalSql = finalSql + "\n" + sql.GenerateAclRoleInsert(usernames, warehouseCode, consts.ConsoleEnv)
	}

	return finalSql
}

func GenerateUsers(list users.UsersList, warehouseCode, phoneNumber string) []users.AclUser {
	var aclUser users.AclUser
	var aclUsers []users.AclUser

	for i := 1; i <= list.Pkr; i++ {
		aclUser.CreateUser(i, warehouseCode, "PKR", phoneNumber)
		aclUsers = append(aclUsers, aclUser)
	}

	for i := 1; i <= list.Crd; i++ {
		aclUser.CreateUser(i, warehouseCode, "CRD", phoneNumber)
		aclUsers = append(aclUsers, aclUser)
	}

	for i := 1; i <= list.Adm; i++ {
		aclUser.CreateUser(i, warehouseCode, "ADM", phoneNumber)
		aclUsers = append(aclUsers, aclUser)
	}

	return aclUsers
}

func GenerateUsersApi(warehouseUsers handlers.WarehouseUsers, warehouseCode string) []users.AclUser {
	var aclUser users.AclUser
	var aclUsers []users.AclUser

	for i := 1; i <= warehouseUsers.Pkr; i++ {
		aclUser.CreateUser(i, warehouseCode, "PKR", warehouseUsers.WarehousePhone)
		aclUsers = append(aclUsers, aclUser)
	}

	for i := 1; i <= warehouseUsers.Crd; i++ {
		aclUser.CreateUser(i, warehouseCode, "CRD", warehouseUsers.WarehousePhone)
		aclUsers = append(aclUsers, aclUser)
	}

	for i := 1; i <= warehouseUsers.Adm; i++ {
		aclUser.CreateUser(i, warehouseCode, "ADM", warehouseUsers.WarehousePhone)
		aclUsers = append(aclUsers, aclUser)
	}

	return aclUsers
}

func PublishAclUsers(users []users.AclUser, warehouseCode string) {
	log.Debugln("Publish users to ACL")

	var rolePickingCode string
	var roleConsoleCode string

	token := strings.Split(files.ReadFile(files.ReadToken(consts.TokenFile)), "\n")[0]

	defer log.Debugln("End of call endpoint")

	for _, user := range users {
		switch user.RoleCode {
		case "PKR":
			rolePickingCode = consts.PickingCodePkr
			roleConsoleCode = consts.ConsoleCodePkr
			break
		case "CRD":
			rolePickingCode = consts.PickingCodeCrd
			roleConsoleCode = consts.ConsoleCodeCrd
			break
		case "ADM":
			rolePickingCode = consts.PickingCodeAdm
			roleConsoleCode = consts.ConsoleCodeAdm
			break
		default:
			log.Fatalln("AuthUser dont identify")
		}

		//TODO call to return token to cas/token endpoint
		//upload users to acl appPicking with role and store code
		userAcl := json.GenerateAclJson(consts.AppPickingEnv, warehouseCode, rolePickingCode, false)
		http.AclEndpointCall(userAcl, user.Username, token)

		//upload users to acl console with role and store code
		userAcl = json.GenerateAclJson(consts.ConsoleEnv, warehouseCode, roleConsoleCode, false)
		http.AclEndpointCall(userAcl, user.Username, token)
	}

}

func PublishAclUsersApi(users []users.AclUser, warehouseCode, token string) error {
	log.Debugln("Publish users to ACL")

	var err error = nil

	var rolePickingCode string
	var roleConsoleCode string

	defer log.Debugln("End of call endpoint")

	for _, user := range users {
		switch user.RoleCode {
		case "PKR":
			rolePickingCode = consts.PickingCodePkr
			roleConsoleCode = consts.ConsoleCodePkr
			break
		case "CRD":
			rolePickingCode = consts.PickingCodeCrd
			roleConsoleCode = consts.ConsoleCodeCrd
			break
		case "ADM":
			rolePickingCode = consts.PickingCodeAdm
			roleConsoleCode = consts.ConsoleCodeAdm
			break
		default:
			log.Fatalln("AuthUser dont identify")
		}

		//TODO call to return token to cas/token endpoint

		//upload users to acl appPicking with role and store code
		userAcl := json.GenerateAclJson(consts.AppPickingEnv, warehouseCode, rolePickingCode, false)
		err = http.AclEndpointCall(userAcl, user.Username, token)
		if err != nil {
			return err
		}
		//upload users to acl console with role and store code
		userAcl = json.GenerateAclJson(consts.ConsoleEnv, warehouseCode, roleConsoleCode, false)
		err = http.AclEndpointCall(userAcl, user.Username, token)
		if err != nil {
			return err
		}
	}

	return nil
}

func PublisDrivershRoles(drivers []string) {
	log.Debugln("Publish to acl the drivers")
	//TODO call to return token to cas/token endpoint
	token := strings.Split(files.ReadFile(files.ReadToken(consts.TokenFile)), "\n")[0]
	for _, driver := range drivers {
		log.Debugf("Publish driver %v", driver)
		driverJson := json.GenerateAclJson(consts.DriverApp, "", consts.DriverRole, true)
		http.AclEndpointCall(driverJson, driver, token)
	}
}
