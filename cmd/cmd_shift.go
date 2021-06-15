package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg"
	"github.com/spf13/cobra"
	"log"
)

var (
	ShiftCmd = &cobra.Command{
		Use:     "shift [COLUMN_NAME] [NUMBER]",
		Aliases: []string{"sf"},

		Short: "Shift column value",
		Long:  "Shift column",
		Example: `
shift close and 2
  siga sf close 2 < sample.tsv
`,
		Run: runCommandShift,
	}
)

func init() {
	ShiftCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandShift(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("More than two COLUMN_NAME or NUMBER should be declared")
	}

	df := pkg.ReadDataFrameByStdinTsv()
	columnName := args[0]
	offset := pkg.ParseInt(args[1])

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	vector = vector.Shift(offset)

	if label == "" {
		label = fmt.Sprintf("shift_%s_%d", columnName, offset)
	}
	err = df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
