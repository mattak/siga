package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
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
	LessEqualCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandLessEqual(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	creator := pipeline.CobraCommandInput{cmd, args}.CreateLessEqualCommandOption(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.CreatePipe(df).Execute()
	df.PrintTsv(IsPreciseOutput)
}
