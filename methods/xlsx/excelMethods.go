package xlsx

import (
	"drivers-create/methods/log"
	"drivers-create/methods/sql"
	"fmt"

	"github.com/tealeg/xlsx"
)

func readExcelFile() {
	excelFileName := "ejemplo.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %s", err)
	}

	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)

		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text)
			}
			fmt.Println("")
		}
	}
}

func ExpeditionLayout(excelFileName string) (string, error) {
	log.Debugln("Creation Expedition Layout sql file")
	var err error
	fQuery := ""

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Errorf("Error reading expedition layout, Error: %v", err)
		return "", err
	}

	layout := xlFile.Sheets[0]
	for i, row := range layout.Rows {
		if i > 0 {
			warehouseCode := row.Cells[0].String()
			typeE := row.Cells[1].String()
			locationZone := row.Cells[2].String()
			area := row.Cells[3].String()
			position := row.Cells[4].String()
			templateArea := row.Cells[5].String()
			templatePosition := row.Cells[6].String()
			shippingSecuence := row.Cells[7].String()
			priority := row.Cells[8].String()
			closer_sorter := row.Cells[9].String()
			active := row.Cells[10].String()
			locationTemplate := row.Cells[11].String()
			location := row.Cells[12].String()

			if len(position) < 2 {
				position = "0" + position
			}

			query := sql.GenerateExpeditionLayoutSql(warehouseCode, typeE, locationZone, area, position, templateArea, templatePosition, shippingSecuence, priority, closer_sorter, active, locationTemplate, location)

			fQuery = fQuery + "\n" + query
		}
	}

	log.Debugln("Finish generating file")
	return fQuery, err
}

func PickingLayout(excelFileName string) (string, error) {
	var err error

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Errorf("Error reading expedition layout, Error: %v", err)
		return "", err
	}

	layout := xlFile.Sheets[0]
	for _, row := range layout.Rows {
		for _, cell := range row.Cells {
			fmt.Println(cell)
		}
	}

	return "", err
}
