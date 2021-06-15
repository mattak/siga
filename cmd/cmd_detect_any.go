package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg"
	"github.com/spf13/cobra"
	"log"
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
	df := pkg.ReadDataFrameByStdinTsv()
	if len(args) < 1 {
		log.Fatal("COLUMN_NAME must be declared")
	}
	columnName := args[0]
	detectValue := 1.0
	if len(args) >= 2 {
		detectValue = pkg.ParseFloat64(args[1])
	}

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	result := vector.HasAnyValue(detectValue)
	fmt.Println(result)
}
