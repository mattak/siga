package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	GreaterThanCmd = &cobra.Command{
		Use:     "greater_than [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"gt"},
		Short: "Greater than comparison",
		Long: "Greater than comparison",
		Example: `
  siga gt column1 100
  siga gt column1 column2
`,
		Run: runCommandGreaterThan,
	}
)

func init() {
	GreaterThanCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandGreaterThan(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(args)

	vector := matrix.GreaterThan()
	if label == "" {
		label = fmt.Sprintf("gt_%s_%s", args[0], args[1])
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
