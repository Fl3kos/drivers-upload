package file

import (
	"bufio"
	"encoding/csv"
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

func CreationFileRoute(route, extension string) string {
	fileName := fmt.Sprint("./files/", route, "-", common.GetDate(), ".", extension)

	return fileName
}

func ReadFileRoute(route, extension string) string {
	fileName := fmt.Sprint("./filesToRead/", route, ".", extension)

	return fileName
}
