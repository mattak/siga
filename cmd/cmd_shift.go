package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
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
	outputOption := pipeline.OutputOption{ColumnName: label}
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateShiftPipe(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = pipe.Execute(df)
	df.PrintTsv(IsPreciseOutput)
}
