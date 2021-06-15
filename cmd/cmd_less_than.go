package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	LessThanCmd = &cobra.Command{
		Use:     "less_than [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"lt"},
		Short:   "Less than comparison",
		Long:    "Less than comparison",
		Example: `
  siga lt column1 100
  siga lt column1 column2
`,
		Run: runCommandLessThan,
	}
)

func init() {
	LessThanCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandLessThan(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	creator := pipeline.CobraCommandInput{cmd, args}.CreateLessThanCommandOption(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.CreatePipe(df).Execute()
	df.PrintTsv(IsPreciseOutput)
}
