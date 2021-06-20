package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	SelectLastValueCmd = &cobra.Command{
		Use:     "select_last_value [COLUMN_NAME]+",
		Aliases: []string{"slv"},

		Short: "select record",
		Long:  "Select record",
		Example: `
  siga slv value < sample.tsv
`,
		Run: runCommandSelectLastValue,
	}
)

func init() {
}

func runCommandSelectLastValue(cmd *cobra.Command, args []string) {
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateSelectLastValuePipeFloat64()
	df := dataframe.ReadDataFrameByStdinTsv()
	fmt.Printf("%.3f\n", pipe.Execute(df))
}
