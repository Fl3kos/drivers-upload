package dni

import (
	mt "drivers-create/methods"
	"fmt"
	"strconv"
)

func ComprobeAllDnis(allDnis []string) bool {
	_continue := true

	for _, dni := range allDnis {
		fmt.Println("Coprobando dni:", dni)
		var letter = dni[8:9]
		if !mt.IsNumber(dni[:1]) {
			fmt.Println(dni)
			break
		}
		correctLetter := calculateTheLetterOfDni(dni)
		if letter == correctLetter {
			fmt.Println("DNI correcto")
			fmt.Println(dni, correctLetter)
			_continue = true
		} else {
			fmt.Println("DNI incorrecto")
			fmt.Println(dni, correctLetter)
			_continue = false
		}
	}

	return _continue
}

func calculateTheLetterOfDni(dni string) string {
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var dniNumber = dni[:8]
	var dniNumberInt, _ = strconv.Atoi(dniNumber)
	dniLetter := letters[dniNumberInt%23]
	return dniLetter
}
