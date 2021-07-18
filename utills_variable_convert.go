package easysdk

import (
	"strconv"
)

func StringToUint64(text string) (uint64, error) {
	u64, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return 0, err
	}
	return u64, nil
}

func StringToUint(text string) (uint, error) {
	u64, err := strconv.ParseUint(text, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(u64), nil
}

func StringToFloat64(text string) (float64, error) {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func StringToFloat32(text string) (float32, error) {
	value, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}
