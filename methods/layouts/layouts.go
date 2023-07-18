package layouts

import (
	"drivers-create/methods/file"
	"drivers-create/methods/json"
	"drivers-create/methods/log"
	excel "drivers-create/methods/xlsx"
)

func GeneratePickingLayout() {
	log.Debugln("Generate SQL file to Picking Layout")
	pickingQuery, err := excel.PickingLayout("./filesToRead/layouts/picking.xlsx")

	err = file.GenerateFile(pickingQuery, file.CreationFilePickingSql("picking", "sql"))
	if err != nil {
		log.Errorf("Error generating picking sql file, Error: %v", err)
	}
}

func GenerateExpeditionLayout()(string) {
	log.Debugln("Generate SQL file to Expedition Layout")
	expeditionQuery, err, sorterRow, warehouse := excel.ExpeditionLayout("./filesToRead/layouts/expedition.xlsx")

	err = file.GenerateFile(expeditionQuery, file.CreationFileExpeditionSql("expedition", "sql"))
	if err != nil {
		log.Errorf("Error generating expedition sql file, Error: %v", err)
	}

	sorterMap := json.GenerateSorterMap(sorterRow, warehouse)
	return sorterMap
}
