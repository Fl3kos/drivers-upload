package dni

import (
	mt "drivers-create/methods"
	logs "drivers-create/methods/log"
	"strconv"
)

func ComprobeAllDnis(allDnis []string) bool {
	logs.DebugLog.Println("Comprobe Dnis")
	var _continue bool

	for _, dni := range allDnis {
		logs.DebugLog.Println("Comprobing DNI:", dni)
		var letter = dni[8:9]
		if !mt.IsNumber(dni[:1]) {
			logs.WarningLog.Printf("DNI %v is a NIE", dni)
			break
		}
		correctLetter := calculateTheLetterOfDni(dni)
		if letter == correctLetter {
			logs.DebugLog.Printf("DNI %v is correct", dni)
			_continue = true
		} else {
			logs.ErrorLog.Printf("Incorrect DNI %v, the correct letter is %v", dni, correctLetter)
			_continue = false
		}
	}

	return _continue
}

func calculateTheLetterOfDni(dni string) string {
	logs.DebugLog.Println("Calculating the letter")
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var dniNumber = dni[:8]
	var dniNumberInt, _ = strconv.Atoi(dniNumber)
	dniLetter := letters[dniNumberInt%23]
	return dniLetter
}
