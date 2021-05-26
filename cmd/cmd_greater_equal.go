package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	GreaterEqualCmd = &cobra.Command{
		Use:     "greater_equal [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"ge"},
		Short: "Greater equal comparison",
		Long: "Greater equal comparison",
		Example: `
  siga ge column1 100
  siga ge column1 column2
`,
		Run: runCommandGreaterEqual,
	}
)

func init() {
	GreaterEqualCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandGreaterEqual(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(args)

	vector := matrix.GreaterEqual()

	if label == "" {
		label = fmt.Sprintf("ge_%s_%s", args[0], args[1])
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
