package main

import (
	"drivers-create/methods/file"
	"drivers-create/methods/log"
	"drivers-create/methods/layouts"

	//"github.com/tealeg/xlsx"
	"fmt"
)

func main() {
	log.InitLogger()

	var sorterMap string
	for {
		fmt.Printf("Wich layouts are you generate? Picking (p), Expedition (e), All (a)")
		var answer string
		fmt.Scanln(&answer)

		switch answer {
		case "p":
			//generate picking layout
			layouts.GeneratePickingLayout()
			break
		case "e":
			//generate expedition layout
			sorterMap = layouts.GenerateExpeditionLayout()
			break
		case "a":
			//generate expedition and picking layout
			layouts.GeneratePickingLayout()
			sorterMap = layouts.GenerateExpeditionLayout()
			break
		default:
			fmt.Println("Select a validate option")

		}
	}

	if sorterMap != "" {
		err := file.GenerateFile(sorterMap, file.CreationSorterMap())
		if err != nil {
			log.Errorf("Error generating picking sql file, Error: %v", err)
		}
	}

	fmt.Println("Finish")
	//fmt.Println("Picking\n" + pickingQuery)
}

