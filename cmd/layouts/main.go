package main

import (
	"drivers-create/methods/file"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	"drivers-create/methods/xlsx"

	//"github.com/tealeg/xlsx"
	"fmt"
)

func main() {
	log.InitLogger()
	expeditionQuery, err, excel, warehouse := xlsx.ExpeditionLayout("./filesToRead/layouts/expedition.xlsx")
	pickingQuery, err := xlsx.PickingLayout("./filesToRead/layouts/picking.xlsx")
	sorterMap := json.GenerateSorterMap(excel, warehouse)

	err = file.GenerateFile(expeditionQuery, file.CreationFileExpeditionSql("expedition", "sql"))
	if err != nil {
		log.Errorf("Error generating expedition sql file, Error: %v", err)
	}

	err = file.GenerateFile(pickingQuery, file.CreationFilePickingSql("picking", "sql"))
	if err != nil {
		log.Errorf("Error generating picking sql file, Error: %v", err)
	}

	if sorterMap != "" {
		err = file.GenerateFile(sorterMap, file.CreationSorterMap())
		if err != nil {
			log.Errorf("Error generating picking sql file, Error: %v", err)
		}
	}

	fmt.Println("Finish")
	//fmt.Println("Picking\n" + pickingQuery)
}
