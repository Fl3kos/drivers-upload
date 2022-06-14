package main

import (
	"bufio"
	"crypto/sha256"
	dniM "drivers-create/methods/dni"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var allDnis []string
var allUsers []string
var allPasswords []string
var allNames []string
var allPhones []string
var shopCode string

var toDate = exportToDate()

var m int

func main() {
	//readfiles method
	dnisFromFile := readFile("./filesToRead/dnis.txt")

	substring := dnisFromFile[:len(dnisFromFile)-1]
	allDnis = strings.Split(substring, "\n")

	//make trim to all dnis
	for i, dni := range allDnis {
		allDnis[i] = strings.TrimSpace(dni)
	}

	_continue := dniM.ComprobeAllDnis(allDnis)
	//_continue := comprobeAllDnis()

	// convertir los dni en usuarios
	if _continue {
		fmt.Println("allDnis: ", allDnis)
		m = len(allDnis)
		allUsers = convertAllDnisToUsers()
		fmt.Println("allUsers: ", allUsers)

		// convertir las contraseñas si no existen
		allPasswords = convertAllUsersToPasswords()
		fmt.Println("allPasswords: ", allPasswords)

		//fmt.Println("pon los nombres separdos por comas (,)")
		names := readFile("./filesToRead/names.txt")

		phonesNumber := readFile("./filesToRead/phoneNumbers.txt")
		allNames = strings.Split(names, "\n")
		allPhones = strings.Split(phonesNumber, "\n")

		shopCode = strings.Split(readFile("./filesToRead/shopCode.txt"), "\n")[0]

		//make trim to all users
		for i, name := range allNames {
			allNames[i] = strings.TrimSpace(name)
		}

		// creacion de las queries
		json := generateJson()
		sql := generateSql()
		allNames := transformAllNames()
		driversInsert := generateSqlLiteInsertDriversTable()
		relationsInsert := generateSqlLiteInsertRelationTable()
		sqlLiteInserts := driversInsert + "\n\n" + relationsInsert

		//passwords := strings.Join(allPasswords, ",")
		getDate()

		// creacion de los files
		sqlFileName := "./files/insertQuery-" + toDate + ".sql"
		jsonFileName := "./files/usersCouchbase-" + toDate + ".json"
		namesFileName := "./files/names-" + toDate + ".txt"
		usersAndPasswordsFileName := "./files/usersAndPasswords-" + toDate + ".txt"
		insertSqlFileName := "./files/insertSQLIteQuery-" + toDate + ".sql"

		generateFile(json, jsonFileName)
		generateFile(sql, sqlFileName)
		generateFile(allNames, namesFileName)
		generateFile(usersAndPasswords(), usersAndPasswordsFileName)
		generateFile(sqlLiteInserts, insertSqlFileName)
		//WriteCsv()
		// end
	}

	// futuramente hacer un menú
}

//readFiles

func comprobeAllDnis() bool {
	_continue := true

	for _, dni := range allDnis {
		fmt.Println("Coprobando dni:", dni)
		var letter = dni[8:9]
		if !isNumber(dni[:1]) {
			fmt.Println(dni)
			break
		}
		correctLetter := calculateTheLetterOfDni(dni)
		if letter == correctLetter {
			fmt.Println("DNI correcto")
			fmt.Println(dni, correctLetter)
			_continue = true
		} else {
			fmt.Println("DNI incorrecto")
			fmt.Println(dni, correctLetter)
			_continue = false
		}
	}

	return _continue
}

func calculateTheLetterOfDni(dni string) string {
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var dniNumber = dni[:8]
	var dniNumberInt, _ = strconv.Atoi(dniNumber)
	dniLetter := letters[dniNumberInt%23]
	return dniLetter
}

func convertAllDnisToUsers() []string {
	var allUsers = []string{}
	for _, dni := range allDnis {
		var user string
		var letter = dni[8:9]

		if isNumber(dni[0:1]) {
			user = letter + dni[1:8]
		} else {
			user = dni[:1] + letter + dni[2:8]
		}
		user = strings.ToUpper(user)
		allUsers = append(allUsers, user)
	}

	return allUsers
}

//comprobar si un caracter es un numero
func isNumber(c string) bool {
	_, err := strconv.Atoi(c)
	if err == nil {
		return true
	}
	return false
}

func convertAllUsersToPasswords() []string {
	var allPasswords = []string{}
	var letter string
	var password string
	for _, user := range allUsers {
		if isNumber(user[1:2]) {
			letter = strings.ToLower(user[0:1])
		} else {
			letter = strings.ToLower(user[1:2])
		}

		password = user[:7] + letter

		allPasswords = append(allPasswords, password)
	}
	return allPasswords
}

func exportToDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func getDate() {
	var lines []string
	var textSplit string
	for _, line := range strings.Split(toDate, "-") {
		line = strings.TrimSpace(line)

		if line != "" {
			lines = append(lines, line)
		}

		textSplit = textSplit + line
	}

	toDate = textSplit
	fmt.Println(toDate)
	fmt.Println(textSplit)
}

//generate a sql file, import the sql text
func generateFile(text string, fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return

	}

	l, err := f.WriteString(text)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

//read text from a file
func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return ""
	}

	return text
}

//this method generate sql query
//the length of array is the value of x
func generateSql() string {
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

func generateSqlLiteInsertDriversTable() string {
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

func generateSqlLiteInsertRelationTable() string {
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

//encode the password
func encodePassword(password string) string {
	encrypted := sha256.Sum256([]byte(password))
	ps := hex.EncodeToString(encrypted[:])
	return ps
}

//this method format the user and password for the json text
func formatUserAndPassword(user, password string) (userF string, passF string) {
	userF = "\"" + user + "\""
	passF = "\"" + password + "\""
	return
}

//this method generate json text
//the length of arrays is the value of x
func generateJson() string {
	json := "[\n"

	collection := "\"authentication\""
	userType := "\"userType\": \"ECOMMERCE_USER\""
	for i := 0; i < m; i++ {
		name := strings.Split(allNames[i], " ")
		firstname := "\"" + name[0] + "\""
		lastname := "\""

		for j := 1; j < len(name); j++ {
			lastname = lastname + name[j] + " "
		}

		lastname = strings.TrimSpace(lastname)
		lastname = lastname + "\""

		ps := encodePassword(allPasswords[i])

		user, pass := formatUserAndPassword(allUsers[i], ps)

		value := "\t{\n\t\t\"username\" : " + user + " ,\n\t\t\"password\" : " + pass + " ,\n\t\t\"firstname\" : " + firstname + " ,\n\t\t\"lastname\" : " + lastname + " ,\n\t\t\"collection\" : " + collection + ",\n\t\t" + userType + "\n\t}"

		if i != m-1 {
			value = value + ",\n"
		}

		json = json + value
	}

	json = json + "\n]"
	return json
}

func transformAllNames() string {
	var allNamesT string
	for _, name := range allNames {
		var nameF = strings.ToLower(name)
		//remplazar espacios por guiones
		nameF = strings.ReplaceAll(nameF, " ", "-")
		allNamesT = allNamesT + nameF + "\n"
	}
	return allNamesT
}

func usersAndPasswords() string {
	var usersAndPasswords = "NAME\n"
	for i := 0; i < m; i++ {
		usersAndPasswords = usersAndPasswords + allNames[i] + "\n"
	}

	usersAndPasswords = usersAndPasswords + "\nUSERS\n"
	for i := 0; i < m; i++ {
		usersAndPasswords = usersAndPasswords + allUsers[i] + "\n"
	}

	usersAndPasswords = usersAndPasswords + "\nPASSWORDS\n"
	for i := 0; i < m; i++ {
		usersAndPasswords = usersAndPasswords + allPasswords[i] + "\n"
	}

	return usersAndPasswords
}

func WriteCsv() {
	// Crea un archivo
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Escribir BOM UTF-8
	f.WriteString("\xEF\xBB\xBF")
	// Crea una nueva secuencia de archivos de escritura
	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "Liu Bei", "23"},
		{"2", "Zhang Fei", "23"},
		{"3", "Guan Yu", "23"},
		{"4", "Zhao Yun", "23"},
		{"5", "Huang Zhong", "23"},
		{"6", "Ma Chao", "23"},
	}
	//Entrada de datos
	w.WriteAll(data)
	w.Flush()
}
