package file

import (
	"bufio"
	"fmt"
	"os"

	common "drivers-create/methods"
	logs "drivers-create/methods/log"
)

//generate a sql file, import the sql text
func GenerateFile(text string, fileName string) {
	logs.DebugLog.Printf("Generating %v file", fileName)

	f, err := os.Create(fileName)

	if err != nil {
		logs.ErrorLog.Printf("Was an error creating %v file. Error: %v", fileName, err)
		return

	}

	_, err = f.WriteString(text)

	if err != nil {
		logs.ErrorLog.Printf("Error writing %v file. Error: %v", fileName, err)
		f.Close()
		return
	}

	logs.DebugLog.Printf("The file %v was write succesfull", fileName)

	err = f.Close()

	if err != nil {
		logs.ErrorLog.Printf("Error closing file %v. Error %v", fileName, err)
		return
	}
}

//read text from a file
func ReadFile(fileName string) string {
	logs.DebugLog.Printf("Reading %v file", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		logs.ErrorLog.Printf("Error opening %v file. Error: %v", fileName, err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		logs.ErrorLog.Printf("Error reading %v file. Error: %v", fileName, err)
		return ""
	}

	logs.DebugLog.Printf("File %v was writting succesfull", fileName)

	return text
}

func CreationFileRoute(route, extension string) string {
	fileName := fmt.Sprintf("./files/%v-%v.%v", route, common.GetDate(), extension)
	logs.DebugLog.Printf("Read creation file %v", fileName)

	return fileName
}

func ReadFileRoute(route, extension string) string {
	fileName := fmt.Sprintf("./filesToRead/%v.%v", route, extension)

	logs.DebugLog.Printf("Read route file %v", fileName)

	return fileName
}

func ReadSqliteFile(fileName string) string {
	file := fmt.Sprintf("%v.db", fileName)
	logs.DebugLog.Printf("Read creation file %v", file)
	return file
}
