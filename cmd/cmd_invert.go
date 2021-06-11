package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	InvertCmd = &cobra.Command{
		Use:     "invert [COLUMN_NAME]",
		Aliases: []string{"inv"},

		Short: "Invert column value",
		Long:  "Invert column",
		Example: `
invert close and 2
  siga inv close 2 < sample.tsv
`,
		Run: runCommandInvert,
	}
)

func init() {
	InvertCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandInvert(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than two COLUMN_NAME should be declared")
	}

	df := toolkit.ReadDataFrameByStdinTsv()
	columnName := args[0]

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	vector = vector.Invert()

	if label == "" {
		label = fmt.Sprintf("invert_%s", columnName)
	}
	err = df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
