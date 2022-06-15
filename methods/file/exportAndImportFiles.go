package file

import (
	"bufio"
	"fmt"
	"os"

	common "drivers-create/methods"
)

//generate a sql file, import the sql text
func GenerateFile(text string, fileName string) {
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
func ReadFile(fileName string) string {
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

func CreationFileRoute(route, extension string) string {
	fileName := fmt.Sprint("./files/", route, "-", common.GetDate(), ".", extension)

	return fileName
}

func ReadFileRoute(route, extension string) string {
	fileName := fmt.Sprint("./filesToRead/", route, ".", extension)

	return fileName
}
