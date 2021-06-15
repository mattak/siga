package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	GreaterEqualCmd = &cobra.Command{
		Use:     "greater_equal [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"ge"},
		Short: "Greater equal comparison",
		Long: "Greater equal comparison",
		Example: `
  siga ge column1 100
  siga ge column1 column2
`,
		Run: runCommandGreaterEqual,
	}
)

func init() {
	GreaterEqualCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandGreaterEqual(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	creator := pipeline.CobraCommandInput{cmd, args}.CreateGreaterEqualCommandOption(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.CreatePipe(df).Execute()
	df.PrintTsv(IsPreciseOutput)
}
