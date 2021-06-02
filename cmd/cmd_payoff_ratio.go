package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	PayoffRatioCmd = &cobra.Command{
		Use:     "payoff_ratio [COLUMN_NAME]+",
		Aliases: []string{"payoff", "po"},

		Short: "Calculate payoff ratio",
		Long:  "Calculate payoff ratio",
		Example: `
  siga po column1 < sample.tsv
  siga po column1 column2 < sample.tsv
`,
		Run: runCommandPayoffRatio,
	}
)

func runCommandPayoffRatio(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than one COLUMN_NAME should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	result := make([]PayoffResult, len(args))

	for i := 0; i < len(args); i++ {
		vector, err := df.ExtractColumn(args[i])
		if err != nil {
			log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
		}

		result[i] = vector.PayoffRatio()

		if i == 0 {
			result[i].PrintTsvHeader()
		}

		result[i].PrintTsvData(IsPreciseOutput)
	}
}
