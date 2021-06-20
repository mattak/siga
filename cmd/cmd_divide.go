package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	DivideCmd = &cobra.Command{
		Use:     "divide [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"div"},

		Short: "Divide",
		Long:  "Divide",
		Example: `
divide close and 2
  siga div close 2 < sample.tsv
`,
		Run: runCommandDivide,
	}
)

func init() {
	DivideCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandDivide(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateDividePipe(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = pipe.Execute(df)
	df.PrintTsv(IsPreciseOutput)
}
