package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	ConstCmd = &cobra.Command{
		Use:     "const [NUMBER]+",
		Aliases: []string{"c"},

		Short: "Const",
		Long:  "Const",
		Example: `
add const 1
  siga c 1 < sample.tsv
`,
		Run: runCommandConst,
	}
)

func init() {
	ConstCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandConst(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	creator := pipeline.CobraCommandInput{cmd, args}.CreateConstCommandOption(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.CreatePipe(df).Execute()
	df.PrintTsv(IsPreciseOutput)
}
