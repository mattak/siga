package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var (
	DivideCmd = &cobra.Command{
		Use:     "divide [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"div"},

		Short: "Divide",
		Long:  "Divide",
		Example: `
divide close and 2
  siga div close 2 < sample.tsv
`,
		Run: runCommandDivide,
	}
)

func init() {
	DivideCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandDivide(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("More than two COLUMN_NAME or NUMBER should be declared")
	}

	df := dataframe.ReadDataFrameByStdinTsv()

	matrix := make(dataframe.Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = dataframe.CreateVector(len(df.Labels))

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

	vector := matrix.Divide()

	if label == "" {
		label = fmt.Sprintf("divide_%s", strings.Join(args, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
