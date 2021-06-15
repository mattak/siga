package dataframe

import (
	"fmt"
	"strings"
)

func (df *DataFrame) PrintTsv(precise bool) {
	floatFormat := "\t%.3f"
	if precise {
		floatFormat = "\t%f"
	}

	fmt.Println(strings.Join(df.Headers, "\t"))
	for i := 0; i < len(df.Labels); i++ {
		fmt.Print(df.Labels[i])
		for j := 0; j < len(df.Data[i]); j++ {
			fmt.Printf(floatFormat, df.Data[i][j])
		}
		fmt.Println()
	}
}
