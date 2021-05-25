package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	EqualCmd = &cobra.Command{
		Use:     "equal [COLUMN_NAME|NUMBER]+",
		Aliases: []string{"eq", "and"},
		Short:   "Equal comparison",
		Long:    "Equal comparison",
		Example: `
  siga eq column1 100
  siga eq column1 column2
  siga and column1 column2 100
`,
		Run: runCommandEqual,
	}
)

func init() {
}

func runCommandEqual(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(args)

	vector := matrix.Equal()
	label := fmt.Sprintf("eq_%s", strings.Join(args, "_"))
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
