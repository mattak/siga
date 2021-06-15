package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	AnomalyCmd = &cobra.Command{
		Use:     "anomaly [COLUMN_NAME] [SPAN]",
		Aliases: []string{"an"},

		Short: "Anomaly calculation",
		Long:  "Anomaly calculation. output is relative to sigma `(v-mean)/σ`. 1.0 means 1σ",
		Example: `
anomaly calculation by 10 data points:
  siga ad volume 10 < sample.tsv
`,
		Run: runCommandAnomaly,
	}
)

func init() {
	AnomalyCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandAnomaly(cmd *cobra.Command, args []string) {
	cobraInput := pipeline.CobraCommandInput{cmd, args}
	outputOption := pipeline.OutputOption{ColumnName: label}
	inputOption := cobraInput.CreateAnomalyCommandOption(outputOption)

	df := dataframe.ReadDataFrameByStdinTsv()
	input := inputOption.CreatePipe(df)
	df = input.Execute()
	df.PrintTsv(IsPreciseOutput)
}
