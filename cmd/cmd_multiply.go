package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var (
	MultiplyCmd = &cobra.Command{
		Use:     "multiply [COLUMN_NAME|NUMBER]+",
		Aliases: []string{"ml", "mul"},

		Short: "Multiply",
		Long:  "Multiply",
		Example: `
multiply close and 2
  siga ml close 2 < sample.tsv
`,
		Run: runCommandMultiply,
	}
)

func init() {
	MultiplyCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandMultiply(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("More than two COLUMN_NAME or NUMBER should be declared")
	}

	df := toolkit.ReadDataFrameByStdinTsv()

	matrix := make(toolkit.Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = toolkit.CreateVector(len(df.Labels))

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

	vector := matrix.InnerProduct()

	if label == "" {
		label = fmt.Sprintf("multiply_%s", strings.Join(args, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
