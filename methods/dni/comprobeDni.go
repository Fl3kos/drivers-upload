package dni

import (
	"errors"
	"strconv"
	logs "support-utils/methods/log"

	common "support-utils/methods"
)

func ComprobeAllDnis(allDnis []string) ([]string, error) {
	logs.Debugln("Comprobe Dnis")
	var incorrectDnis []string
	var err error = nil
	var correctLetter string

	for _, dni := range allDnis {
		if dni != "" {
			var letter = dni[8:9]

			if common.IsNumber(dni[:1]) {
				logs.Warningln("Comprobing NIE:", dni)
				correctLetter = calculateTheLetterOfDni(dni)
			} else {
				logs.Debugln("Comprobing DNI:", dni)
				correctLetter = calculateTheLetterOfNie(dni)
			}

			if letter == correctLetter {
				logs.Debugf("DNI %v is correct", dni)
			} else {
				logs.Errorf("Incorrect DNI %v, the correct letter is %v", dni, correctLetter)
				err = errors.New("Has one or more DNIs incorrect")
				incorrectDnis = append(incorrectDnis, dni)
			}
		}
	}

	return incorrectDnis, err
}

func calculateTheLetterOfDni(dni string) string {
	logs.Debugln("Calculating the letter of DNI")
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var dniNumber = dni[:8]
	var dniNumberInt, _ = strconv.Atoi(dniNumber)
	dniLetter := letters[dniNumberInt%23]
	return dniLetter
}

func calculateTheLetterOfNie(nie string) string {
	logs.Debugln("Calculating the letter of NIE")
	var letters = []string{"T", "R", "W", "A", "G", "M", "Y", "F", "P", "D", "X", "B", "N", "J", "Z", "S", "Q", "V", "H", "L", "C", "K", "E"}
	var nieNumber = nie[1:8]

	nieLetter := nie[:1]
	switch nieLetter {
	case "X":
		nieNumber = "0" + nieNumber
	case "Y":
		nieNumber = "1" + nieNumber
	default:
		nieNumber = "2" + nieNumber
	}

	var nieNumberInt, _ = strconv.Atoi(nieNumber)
	correctLetter := letters[nieNumberInt%23]
	return correctLetter
}
