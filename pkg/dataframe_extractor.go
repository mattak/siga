package pkg

import (
	"log"
	"strconv"
)

func (df *DataFrame) ExtractMatrixByColumnNameOrValue(args []string) Matrix {
	matrix := make(Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = CreateVector(len(df.Labels))

		vector, err := df.ExtractColumn(args[i])
		if err == nil {
			matrix[i] = vector
		} else {
			v, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
			} else {
				matrix[i].Fill(v)
			}
		}
	}
	return matrix
}
