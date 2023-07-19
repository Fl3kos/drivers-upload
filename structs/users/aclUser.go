package users

import (
	"fmt"
	"strconv"
	"strings"
	"support-utils/methods/log"
	numtoletter "support-utils/methods/numToLetter"
)

type AclUser struct {
	AuthUser
	RoleCode string
}

func (u *AclUser) CreateUser(i int, warehouseCode string, userType string, phoneNumber string) {
	number := strconv.Itoa(i)

	for j := len(number); j < 3; j++ {
		number = "0" + number
	}

	for i := len(warehouseCode); i < 5; i++ {
		warehouseCode = "0" + warehouseCode
	}

	userAndPassword := fmt.Sprintf("%v%v%v", userType, warehouseCode, number)

	u.Username = userAndPassword
	u.Password = userAndPassword

	switch userType {
	case "PKR":
		u.Firstname = "APP PICKING"
		u.RoleCode = string(userType)
		break
	case "CRD":
		u.Firstname = "COORDINADOR"
		u.RoleCode = string(userType)
		break
	case "ADM":
		u.Firstname = "ADMINISTRADOR"
		u.RoleCode = string(userType)
		break
	default:
		log.Errorln("AuthUser dont identify")
	}

	lastname, err := numtoletter.IntLetra(i)
	if err != nil {

	}
	u.Lastname = strings.ToUpper(lastname)
	u.Email = fmt.Sprintf("t%ves@stores.diagroup.com", warehouseCode)
	u.Phone = phoneNumber
}
