package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	NormalizeCmd = &cobra.Command{
		Use:     "normalize [COLUMN_NAME]+",
		Aliases: []string{"normal"},
		Short:   "Normalize value",
		Long:    "Normalize value",
		Example: `
normalize value by start value of close
  siga normal close < sample.tsv
`,
		Run: runCommandNormalize,
	}
)

func init() {
	NormalizeCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandNormalize(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than one COLUMN_NAME should be declared")
	}

	df := toolkit.ReadDataFrameByStdinTsv()

	matrix := make(toolkit.Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = toolkit.CreateVector(len(df.Labels))

		column, err := df.ExtractColumn(args[i])
		if err != nil {
			log.Fatalf("COLUMN_NAME required: %s", args[i])
		}
		vector := column.NormalizeByStart()

		newLabel := label
		if label == "" {
			newLabel = fmt.Sprintf("normalize_%s", args[i])
		}

		err = df.AddColumn(newLabel, vector)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
