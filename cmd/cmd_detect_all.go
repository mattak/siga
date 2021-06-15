package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
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
	df := dataframe.ReadDataFrameByStdinTsv()
	if len(args) < 1 {
		log.Fatal("COLUMN_NAME must be declared")
	}
	columnName := args[0]
	detectValue := 1.0
	if len(args) >= 2 {
		detectValue = util.ParseFloat64(args[1])
	}

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	result := vector.IsAllValue(detectValue)
	fmt.Println(result)
}
