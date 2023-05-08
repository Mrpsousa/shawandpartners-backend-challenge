package utils

import "strconv"

func IsNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
