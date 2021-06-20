package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	DeviationsCmd = &cobra.Command{
		Use:     "deviations [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"dev", "stdev"},

		Short: "Deviation",
		Long:  "Deviation",
		Example: `
deviation by span 5
  siga dev close 5 < sample.tsv
deviation by span 5, 10
  siga dev close 5 10 < sample.tsv
`,
		Run: runCommandDeviations,
	}
)

func init() {
	DeviationsCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandDeviations(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateDeviationsPipe(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = pipe.Execute(df)
	df.PrintTsv(IsPreciseOutput)
}
