package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			fmt.Println("Error parsing price:", err)
			return nil, err
		}
		floats = append(floats, floatValue)
	}
	return floats, nil
}
