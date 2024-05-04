package util

import (
	"strconv"
)

func ConvertToFloat(numStr string) (float64, error) {
	// Convert the amount parameter to a float64
	num, err := strconv.ParseFloat(numStr, 64)
	return num, err
}

func ConvertToInt(numStr string) (int, error) {
	// Convert the amount parameter to a float64
	num, err := strconv.ParseInt(numStr, 10, 8)
	return int(num), err
}
