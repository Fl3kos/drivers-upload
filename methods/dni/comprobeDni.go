package dni

import (
	logs "drivers-create/methods/log"
	"errors"
	"strconv"
)

func ComprobeAllDnis(allDnis []string) ([]string, error) {
	logs.DebugLog.Println("Comprobe Dnis")
	var incorrectDnis []string
	var err error = nil
	//var _continue bool

	for _, dni := range allDnis {
		if dni != "" {
			if dni[:1] == "X" || dni[:1] == "Y" {
				// calculate the letter of nie
				logs.WarningLog.Printf("DNI %v is a NIE", dni)

				var letter = dni[8:9]

				correctLetter := calculateTheLetterOfNie(dni)
				if letter == correctLetter {
					logs.DebugLog.Printf("DNI %v is correct", dni)
				} else {
					logs.ErrorLog.Printf("Incorrect DNI %v, the correct letter is %v", dni, correctLetter)
					err = errors.New("Has one or more DNIs incorrect")
					incorrectDnis = append(incorrectDnis, dni)
				}
			} else {
				logs.DebugLog.Println("Comprobing DNI:", dni)

				var letter = dni[8:9]

				correctLetter := calculateTheLetterOfDni(dni)
				if letter == correctLetter {
					logs.DebugLog.Printf("DNI %v is correct", dni)
					//_continue = true
				} else {
					logs.ErrorLog.Printf("Incorrect DNI %v, the correct letter is %v", dni, correctLetter)
					err = errors.New("Has one or more DNIs incorrect")
					incorrectDnis = append(incorrectDnis, dni)
					//_continue = false
				}
			}
		}
	}

	return incorrectDnis, err
}

func calculateTheLetterOfDni(dni string) string {
	logs.DebugLog.Println("Calculating the letter of DNI")
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var dniNumber = dni[:8]
	var dniNumberInt, _ = strconv.Atoi(dniNumber)
	dniLetter := letters[dniNumberInt%23]
	return dniLetter
}

func calculateTheLetterOfNie(nie string) string {
	logs.DebugLog.Println("Calculating the letter of NIE")
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var nieNumber = nie[1:8]
	if nie[:1] == "X" {
		nieNumber = "0" + nieNumber
	} else {
		nieNumber = "1" + nieNumber
	}
	var nieNumberInt, _ = strconv.Atoi(nieNumber)
	nieLetter := letters[nieNumberInt%23]
	return nieLetter
}
