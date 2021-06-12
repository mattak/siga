package toolkit

import (
	"fmt"
)

func (data Vector) PrintTsv(precise bool) {
	startFloatFormat := "%.3f"
	floatFormat := "\t%.3f"

	if precise {
		startFloatFormat = "%f"
		floatFormat = "\t%f"
	}

	if len(data) > 0 {
		fmt.Printf(startFloatFormat, data[0])
	}

	for i := 1; i < len(data); i++ {
		fmt.Printf(floatFormat, data[i])
	}

	if len(data) > 0 {
		fmt.Println()
	}
}
