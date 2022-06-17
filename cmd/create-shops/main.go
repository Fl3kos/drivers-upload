package main

import (
	files "drivers-create/methods/file"
	"drivers-create/methods/log"
	"drivers-create/methods/sql"
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

		allShopCodes = append(allShopCodes, shopCode)
		allShopNames = append(allShopNames, shopName)
	}

	sqlT := sql.GenerateSqlLiteInsertShopTable(allShopCodes, allShopNames)
	files.GenerateFile(sqlT, files.CreationFileRoute("insertSqlTable", "sql"))
}
