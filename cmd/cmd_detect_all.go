package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	DetectAllCmd = &cobra.Command{
		Use:     "detect_all [COLUMN_NAME] [VALUE]?",
		Aliases: []string{"all"},

		Short: "Detect all value equals VALUE",
		Long:  "Detect all value equals VALUE. default VALUE is 1",
		Example: `
  siga all column1 < sample.tsv

e.g. true data condition
  date column1
  2020-01-01 1
  2020-01-01 1

e.g. false data condition
  date column1
  2020-01-01 1
  2020-01-01 0
`,
		Run: runCommandDetectAll,
	}
)

func init() {
}

func runCommandDetectAll(cmd *cobra.Command, args []string) {
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateDetectAllPipeBool()
	df := dataframe.ReadDataFrameByStdinTsv()
	result := pipe.Execute(df)
	fmt.Println(result)
}
