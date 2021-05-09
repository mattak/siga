package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	LessThanCmd = &cobra.Command{
		Use:     "less_than [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"lt"},
		Short: "Less than comparison",
		Long: "Less than comparison",
		Example: `
  siga lt column1 100
  siga lt column1 column2
`,
		Run: runCommandLessThan,
	}
)

func init() {
}

func runCommandLessThan(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(args)

	vector := matrix.LessThan()
	label := fmt.Sprintf("lt_%s_%s", args[0], args[1])
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
