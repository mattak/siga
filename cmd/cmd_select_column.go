package cmd

import (
	"github.com/mattak/siga/pkg"
	"github.com/spf13/cobra"
	"log"
)

var (
	SelectColumnCmd = &cobra.Command{
		Use:     "select_column [COLUMN_NAME]+",
		Aliases: []string{"sc"},

		Short: "select column",
		Long:  "Select column",
		Example: `
  siga sc open close volume < sample.tsv
`,
		Run: runCommandSelectColumn,
	}
)

func init() {
}

func runCommandSelectColumn(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("COLUMN_NAME should be declared")
	}

	df := pkg.ReadDataFrameByStdinTsv()
	err := df.SelectColumn(args...)
	if err != nil {
		log.Fatal(err)
	}
	df.PrintTsv(IsPreciseOutput)
}
