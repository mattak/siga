package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	DetectAnyCmd = &cobra.Command{
		Use:     "detect_any [COLUMN_NAME] [VALUE]?",
		Aliases: []string{"any"},

		Short: "Detect any value equals VALUE",
		Long:  "Detect any value equals VALUE. default VALUE is 1",
		Example: `
  siga any column1 < sample.tsv

e.g. true data condition
  date column1
  2020-01-01 1
  2020-01-01 0

e.g. false data condition
  date column1
  2020-01-01 0
  2020-01-01 0
`,
		Run: runCommandDetectAny,
	}
)

func init() {
}

func runCommandDetectAny(cmd *cobra.Command, args []string) {
	creator := pipeline.CobraCommandInput{cmd, args}.CreateDetectAnyCommandOption()
	df := dataframe.ReadDataFrameByStdinTsv()
	result := creator.CreatePipeBool(df).Execute()
	fmt.Println(result)
}
