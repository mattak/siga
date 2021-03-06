package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	AddCmd = &cobra.Command{
		Use:     "add [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]+",
		Aliases: []string{},
		Short:   "Add",
		Long:    "Add",
		Example: `
  siga add column1 column2
  siga add column1 5
`,
		Run: runCommandAdd,
	}
)

func init() {
	AddCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandAdd(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateAddPipe(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = pipe.Execute(df)
	df.PrintTsv(IsPreciseOutput)
}
