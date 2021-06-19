package dataframe

import (
	"fmt"
	"strings"
)

func (df *DataFrame) ToTsvString(precise bool) string {
	floatFormat := "\t%.3f"
	if precise {
		floatFormat = "\t%f"
	}

	result := fmt.Sprintln(strings.Join(df.Headers, "\t"))
	for i := 0; i < len(df.Labels); i++ {
		result += fmt.Sprint(df.Labels[i])
		for j := 0; j < len(df.Data[i]); j++ {
			result += fmt.Sprintf(floatFormat, df.Data[i][j])
		}
		result += fmt.Sprintln()
	}
	return result
}

func (df *DataFrame) PrintTsv(precise bool) {
	fmt.Print(df.ToTsvString(precise))
}
