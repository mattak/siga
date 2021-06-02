package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var (
	TradingEvaluationCmd = &cobra.Command{
		Use:     "trading_evaluation [BUY_COLUMN|BUY_NUMBER] [SELL_COLUMN|SELL_NUMBER]",
		Aliases: []string{"te"},
		Short:   "Show trading evaluation",
		Long:    "Show trading evaluation",
		Example: `
  siga te open close < sample.tsv
  siga te open 1.0 < sample.tsv
`,
		Run: runCommandTradingEvaluation,
	}
	isDetailOutput = false
)

func init() {
	TradingEvaluationCmd.Flags().BoolVarP(&isDetailOutput, "detail", "d", false, "output detail result")
}

func runCommandTradingEvaluation(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("BUY_COLUMN_NAME|BUY_NUMBER, SELL_COLUMN_NAME|SELL_NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	buysell := make([]Vector, 2)

	for i := 0; i < 2; i++ {
		vector, err := df.ExtractColumn(args[i])
		if err == nil {
			buysell[i] = vector
		} else {
			v, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
			} else {
				buysell[i] = CreateVector(len(df.Labels))
				buysell[i].Fill(v)
			}
		}
	}

	result, err := CreateTradingEvaluationResult(buysell[0], buysell[1])
	if err != nil {
		log.Fatal(err)
	}
	result.PrintTsvHeader(isDetailOutput)
	result.PrintTsvBody(isDetailOutput, IsPreciseOutput)
}
