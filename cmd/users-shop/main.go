package main

import (
	"drivers-create/methods/acl"
	files "drivers-create/methods/file"
	"drivers-create/methods/gets/getPhones"
	"drivers-create/methods/gets/getShops"
	"drivers-create/methods/http"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	"drivers-create/structs/users"

	"fmt"
	"strings"
)

func main() {
	log.InitLogger()
	list := users.ReturnList()
	warehouses := strings.Split(files.ReadFile(files.ReadFileRoute("shops", "txt")), "\n")
	warehouses = warehouses[:len(warehouses)-1]

	warehouseCodes, _ := getShops.GetShopCodesAndShopNames(warehouses)
	warehouseCode := warehouseCodes[0]

	phone := getPhones.GetAllPhones()[0]

	for {
		fmt.Printf("Are you secured to publish aclUsers with warehouse code %v and phone %v? (y/n) ", warehouseCode, phone)
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			break
		}

		if answer == "n" {
			fmt.Printf("The warehouseCode code or phone not the expected")
			log.Fatalf("The warehouseCode code or phone not the expected")
		}
	}

	aclUsers := acl.GenerateUsers(list, warehouseCode, phone)
	finalJson := json.GenerateUsersJson(aclUsers)

	err := files.GenerateFile(finalJson, files.CreationFileUserList())

	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	// Publish user to auth
	http.AuthEndpointCall(finalJson)

	// publish roles to acl
	for {
		fmt.Printf("Are you publish roles to aclUsers with warehouse code %v? (y/n) ", warehouseCode)
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			acl.PublishAclUsers(aclUsers, warehouseCode)
			break
		}

		if answer == "n" {
			log.Debugln("Roles not publish")
			break
		}
	}

	finalSql := acl.GenerateSql(aclUsers, warehouseCode)
	err = files.GenerateFile(finalSql, files.CreationFileRouteAclSql("ACL", "sql"))
	if err != nil {
		log.Errorf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/logs.log")
	}

	fmt.Println("Finish")
}
