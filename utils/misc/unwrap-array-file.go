package misc

import (
	"strconv"
	"strings"
)

func UnwrapArrayFile(data []byte) []int {
	stringNumber := strings.Split(string(data), ",")
	intArray := []int{}
	for i := 0; i < len(stringNumber); i++ {
		num, _ := strconv.Atoi(stringNumber[i])
		intArray = append(intArray, num)
	}
	return intArray
}
