package sql

import "fmt"

func GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones []string) string {
	query := "INSERT INTO drivers (DNI, UserName, Name, PhoneNumber)\nVALUES "
	value := "( '%v', '%v', '%v', '%v')"

	for i, _ := range allDnis {
		if allDnis[i] != "" {
			valueF := fmt.Sprintf(value, allDnis[i], allUsers[i], allNames[i], allPhones[i])

			if i != len(allUsers)-1 {
				valueF = valueF + ",\n"
			}

			query = query + valueF
		}

	}

	query = query + ";"

	return query
}

func GenerateSqlLiteInsertRelationTable(allDnis []string, shopCodes []string) string {

	query := "INSERT INTO DriversShop (DNI, ShopCode)\nVALUES"
	value := `( '%v', '%v')`

	dnisValue := 0
	for i := 0; i < len(shopCodes); i++ {
		for j := dnisValue; j < len(allDnis); j++ {
			if allDnis[j] == "" {
				dnisValue = j + 1
				break
			}
			valueF := fmt.Sprintf(value, allDnis[j], shopCodes[i])

			if j != len(allDnis)-1 {
				valueF = valueF + ",\n"
			}

			query = query + valueF
		}
	}

	query = query + ";"

	return query
}

func GenerateSqlLiteInsertShopTable(shopCodes, shopNames []string) string {
	query := "INSERT INTO Shop (ShopCode, ShopName)\nVALUES "
	value := "('%v', '%v')"

	for i, _ := range shopCodes {
		valueF := fmt.Sprintf(value, shopCodes[i], shopNames[i])
		fmt.Println(valueF)
		if i != len(shopCodes)-1 {
			valueF = valueF + ",\n"
		}

		query = query + valueF
	}

	query = query + ";"

	return query
}
