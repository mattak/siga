package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var (
	GreaterEqualCmd = &cobra.Command{
		Use:     "greater_equal [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"ge"},
		Short: "Greater equal comparation",
		Long: "Greater equal comparation",
		Example: `
  siga ge column1 100
  siga ge column1 column2
`,
		Run: runCommandGreaterEqual,
	}
)

func init() {
}

func runCommandGreaterEqual(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()

	matrix := make(Matrix, len(args))
	for i := 0; i < 2; i++ {
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

	vector := matrix.GreaterEqual()
	label := fmt.Sprintf("ge_%s_%s", args[0], args[1])
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv()
}
