package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	SimpleMovingAverageCmd = &cobra.Command{
		Use:     "simple_moving_average [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"sma"},

		Short: "simple moving average vector calculation",
		Long:  "simple moving average vector calculation",
		Example: `
  siga sma volume 20
  siga sma volume 20 5 1
`,
		Run: runSimpleMovingAverage,
	}
)

func init() {
	SimpleMovingAverageCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runSimpleMovingAverage(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	// args
	columnName := args[0]

	// dataframe -> vector
	df := dataframe.ReadDataFrameByStdinTsv()
	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	vector.Reverse()

	for i := 1; i < len(args); i++ {
		span := util.ParseInt(args[i])
		if span <= 0 {
			log.Fatalf("SPAN should be more than 1: %d\n", span)
		}

		line := vector.SimpleMovingAverage(span)
		line.Reverse()

		newLabel := label
		if label == "" {
			newLabel = fmt.Sprintf("sma_%s_%d", columnName, span)
		}

		err = df.AddColumn(newLabel, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
