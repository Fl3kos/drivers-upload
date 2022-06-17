package sql

import "fmt"

func GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones []string) string {
	query := "INSERT INTO drivers (DNI, UserName, Name, PhoneNumber)"
	value := `VALUES ( '%v', '%v', '%v', '%v')`

	for i, _ := range allDnis {
		valueF := fmt.Sprintf(value, allDnis[i], allUsers[i], allNames[i], allPhones[i])

		if i != len(allUsers)-1 {
			valueF = valueF + ",\n"
		}

		query = query + valueF
	}

	query = query + ";"

	return query
}

func GenerateSqlLiteInsertRelationTable(allDnis []string, shopCode string) string {

	query := "INSERT INTO DriversShop (DNI, ShopCode)"
	value := `VALUES ( '%v', '%v')`

	for i, dni := range allDnis {
		valueF := fmt.Sprintf(value, dni, shopCode)

		if i != len(allDnis)-1 {
			valueF = valueF + ",\n"
		}

		query = query + valueF
	}

	query = query + ";"

	return query
}
