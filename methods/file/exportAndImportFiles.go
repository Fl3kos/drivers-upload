package file

import (
	"bufio"
	"fmt"
	"os"
	"support-utils/consts"

	common "support-utils/methods"
	logs "support-utils/methods/log"
)

// generate a sql file, import the sql text
func GenerateFile(text string, fileName string) error {
	logs.Debugf("Generating %v file", fileName)

	f, err := os.Create(fileName)

	if err != nil {
		logs.Errorf("Was an error creating %v file. Error: %v", fileName, err)
		return err

	}

	_, err = f.WriteString(text)

	if err != nil {
		logs.Errorf("Error writing %v file. Error: %v", fileName, err)
		f.Close()
		return err
	}

	logs.Debugf("The file %v was write succesfull", fileName)

	err = f.Close()

	if err != nil {
		logs.Errorf("Error closing file %v. Error %v", fileName, err)
		return err
	}

	return err
}

// read text from a file
func ReadFile(fileName string) string {
	logs.Debugf("Reading %v file", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		logs.Errorf("Error opening %v file. Error: %v", fileName, err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		logs.Errorf("Error reading %v file. Error: %v", fileName, err)
		return ""
	}

	logs.Debugf("File %v was writting succesfull", fileName)

	return text
}

func CreationFileRoute(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v-%v.%v", consts.FilesRoute, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationFileRouteSql(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesSQL, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}
func CreationFileRouteSqlShop(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesShopsSQL, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationFileRouteAclSql(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesAclSql, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationFileRouteNames(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesNames, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}
func CreationFileRouteCsv(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesCsv, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}
func CreationFileRouteJson(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesJson, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationFileRouteAclJson(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesAclJson, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationFileUserList() string {
	filePath := fmt.Sprintf("%v/%v/%v-%v.json", consts.FilesRoute, consts.FilesUserListJson, consts.UserListFileName, common.GetDate())
	logs.Debugf("Read creation file %v", filePath)

	return filePath
}

func ReadFileRoute(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v.%v", consts.ReadFileRoute, route, extension)

	logs.Debugf("Read route file %v", fileName)

	return fileName
}

func ReadToken(filename string) string {
	fileName := fmt.Sprintf("%v/%v", consts.ReadFileRoute, filename)

	logs.Debugf("Read token file %v", fileName)

	return fileName
}

func ReadSqliteFile(fileName string) string {
	file := fmt.Sprintf("%v.db", fileName)
	logs.Debugf("Read creation file %v", file)
	return file
}

func ReadTestFile(fileName string) string {
	file := fmt.Sprintf("%v/%v", consts.TestFileRoute, fileName)
	logs.Debugf("Read creation file %v", file)
	return file
}

func ReadUserListFile() string {
	file := fmt.Sprintf("%v/%v.json", consts.ReadFileRoute, consts.UserListFileName)
	logs.Debugf("Read UserList file %v", file)
	return file
}

func CreationFileExpeditionSql(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesExpedition, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationFilePickingSql(route, extension string) string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesPicking, route, common.GetDate(), extension)
	logs.Debugf("Read creation file %v", fileName)

	return fileName
}

func CreationSorterMap() string {
	fileName := fmt.Sprintf("%v/%v/%v-%v.%v", consts.FilesRoute, consts.FilesSorterMap, "sorterMap", common.GetDate(), "json")
	logs.Debugf("creation sorter map in file: %v", fileName)

	return fileName
}
