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
	fQuery := ""

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Errorf("Error reading expedition layout, Error: %v", err)
		return "", err
	}

	layout := xlFile.Sheets[0]
	for i, row := range layout.Rows {
		if i > 0 {
			warehouse_code := row.Cells[0].String()
			typeP := row.Cells[1].String()
			location_format := row.Cells[2].String()
			template := row.Cells[3].String()
			corridor := row.Cells[4].String()
			module := row.Cells[5].String()
			shelf := row.Cells[6].String()
			gap := row.Cells[7].String()
			location_zone := row.Cells[8].String()
			picking_zone := row.Cells[9].String()
			weight := row.Cells[10].String()
			height := row.Cells[11].String()
			width := row.Cells[12].String()
			length := row.Cells[13].String()
			capacity_fee := row.Cells[14].String()
			restocking_fee := row.Cells[15].String()
			picking_sequence := row.Cells[16].String()
			putaway_sequence := row.Cells[17].String()
			direction := row.Cells[18].String()
			blocked := row.Cells[19].String()
			active := row.Cells[20].String()
			rotation := row.Cells[21].String()
			cycle_count := row.Cells[22].String()
			location_template := row.Cells[23].String()
			location := row.Cells[24].String()

			if len(corridor) < 2 {
				corridor = "0" + corridor
			}
			if len(module) < 2 {
				module = "0" + module
			}
			if len(shelf) < 2 {
				shelf = "0" + shelf
			}
			if len(gap) < 2 {
				gap = "0" + gap
			}

			query := sql.GeneratePickingLayoutSql(warehouse_code, typeP, location_format, template, corridor, module, shelf, gap, location_zone, picking_zone, weight, height, width, length, capacity_fee, restocking_fee, picking_sequence, putaway_sequence, direction, blocked, active, rotation, cycle_count, location_template, location)

			fQuery = fQuery + "\n" + query
		}
	}

	return fQuery, err
}
