package main

import (
	"drivers-create/methods/file"
	"drivers-create/methods/log"
	"drivers-create/methods/xlsx"
	"fmt"
)

func main() {
	log.InitLogger()
	expeditionQuery, _ := xlsx.ExpeditionLayout("./filesToRead/layouts/expedition.xlsx")
	//pickingQuery, _ := xlsx.PickingLayout("./filesToRead/layouts/picking.xlsx")

	err := file.GenerateFile(expeditionQuery, file.CreationFileExpeditionSql("expedition", "sql"))
	if err != nil {
		log.Errorf("Error generating expedition sql file, Error: %v", err)
	}

	fmt.Println("Finish")
	//fmt.Println("Picking\n" + pickingQuery)
}
