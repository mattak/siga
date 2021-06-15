package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var (
	SumCmd = &cobra.Command{
		Use:     "sum [COLUMN_NAME]+",
		Aliases: []string{},
		Short:   "Sum labels",
		Long:    "Sum labels",
		Example: `
  siga sum column1
`,
		Run: runCommandSum,
	}
)

func init() {
}

func runCommandSum(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than 1 COLUMN_NAME should be declared")
	}
	df := dataframe.ReadDataFrameByStdinTsv()

	sum_vector := make(dataframe.Vector, len(args))

	for i := 0; i<len(args); i++ {
		vector, err := df.ExtractColumn(args[i])
		if err == nil {
			sum_vector[i] = vector.Sum(0, len(vector))
		} else {
			v, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
			} else {
				sum_vector[i] = float64(len(df.Labels)) * v
			}
		}
	}

	sum_vector.PrintTsv(IsPreciseOutput)
}
