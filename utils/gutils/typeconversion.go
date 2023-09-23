package gutils

import "strconv"

// FromStringToInt ...
func FromStringToInt(stringInput string) (int, error) {
	outInt, err := strconv.Atoi(stringInput)
	if err != nil {
		return 0, err
	}
	return outInt, nil
}

func FromStringToFloat64(stringInput string) (float64, error) {
	outFloat64, err := strconv.ParseFloat(stringInput, 64)
	if err != nil {
		return 0, err
	}
	return outFloat64, nil
}
