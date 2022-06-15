package sql

//this method generate sql query
//the length of array is the value of x
func GenerateSql(allUsers []string, m int) string {
	query := "INSERT INTO shipping.DRIVER_GEOLOCATION (DRIVER_ID, LATITUDE, LONGITUDE) \nVALUES "

	for i := 0; i < m; i++ {
		value := "('" + allUsers[i] + "', 0.0, 0.0)"

		if i != m-1 {
			value = value + ",\n"
		}

		query = query + value
	}

	query = query + ";"

	return query
}

func GenerateSqlLiteInsertDriversTable(allUsers, allDnis, allNames, allPhones []string, m int) string {
	query := "INSERT INTO drivers (DNI, UserName, Name, PhoneNumber) \nVALUES "

	for i := 0; i < m; i++ {
		value := "('" + allDnis[i] + "', " + "'" + allUsers[i] + "', '" + allNames[i] + "', '" + allPhones[i] + "')"

		if i != m-1 {
			value = value + ",\n"
		}

		query = query + value
	}

	query = query + ";"

	return query
}

func GenerateSqlLiteInsertRelationTable(allDnis []string, shopCode string, m int) string {
	query := "INSERT INTO DriversShop (DNI, ShopCode) \nVALUES "

	for i := 0; i < m; i++ {
		value := "('" + allDnis[i] + "', '" + shopCode + "')"

		if i != m-1 {
			value = value + ",\n"
		}

		query = query + value
	}

	query = query + ";"

	return query
}
