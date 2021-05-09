package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	LessEqualCmd = &cobra.Command{
		Use:     "less_equal [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"le"},
		Short: "Less equal comparison",
		Long: "Less equal comparison",
		Example: `
  siga le column1 100
  siga le column1 column2
`,
		Run: runCommandLessEqual,
	}
)

func init() {
}

func runCommandLessEqual(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(args)

	vector := matrix.LessEqual()
	label := fmt.Sprintf("le_%s_%s", args[0], args[1])
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv()
}
