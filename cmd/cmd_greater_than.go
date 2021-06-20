package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	GreaterThanCmd = &cobra.Command{
		Use:     "greater_than [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"gt"},
		Short:   "Greater than comparison",
		Long:    "Greater than comparison",
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
	outputOption := pipeline.OutputOption{ColumnName: label}
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateGreaterThanPipe(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = pipe.Execute(df)
	df.PrintTsv(IsPreciseOutput)
}
