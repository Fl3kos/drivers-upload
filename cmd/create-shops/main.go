package main

import (
	"drivers-create/consts"
	files "drivers-create/methods/file"
	"drivers-create/methods/log"
	"drivers-create/methods/sql"
	"drivers-create/methods/sql/sqlite"
	"fmt"
	"strings"
)

func main() {
	log.InitLogger()

	shops := files.ReadFile(files.ReadFileRoute("shops", "txt"))
	shops = shops[:len(shops)-1]

	allShops := strings.Split(shops, "\n")

	var allShopCodes []string
	var allShopNames []string

	for i, _ := range allShops {
		shop := strings.Split(allShops[i], "-")

		shopCode := shop[0]
		shopName := shop[1]

		shopCode = strings.TrimSpace(shopCode)
		shopName = strings.TrimSpace(shopName)

		allShopCodes = append(allShopCodes, shopCode)
		allShopNames = append(allShopNames, shopName)
	}

	sqlT := sql.GenerateSqlLiteInsertShopTable(allShopCodes, allShopNames)
	err := files.GenerateFile(sqlT, files.CreationFileRouteSqlShop("insertSqlTable", "sql"))
	if err != nil {
		log.ErrorLog.Printf("Error generating file: %v", err)
		fmt.Println("Error generating files, check the logs /logs/lo")
	}

	err = sqlite.InsertSqlite(sqlT, consts.SqliteDatabase)

	if err != nil {
		log.ErrorLog.Printf("Error with the insert shop table. Error: %v", err)
		fmt.Println("Was an error has ocurred, please check the logs to solve manualy")
	}

	fmt.Println("Finish")
}
