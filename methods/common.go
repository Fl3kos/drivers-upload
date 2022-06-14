package methods

import "strconv"

func IsNumber(c string) bool {
	_, err := strconv.Atoi(c)
	if err == nil {
		return true
	}
	return false
}
