package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	MartinegaleMovingAverageCmd = &cobra.Command{
		Use:     "martinegale_moving_average [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"mma"},

		Short: "martinegale average vector calculation",
		Long:  "martinegale average vector calculation",
		Example: `
  siga mma volume 20
  siga mma volume 20 5 1
`,
		Run: runMartinegaleMovingAverage,
	}
	martinegaleThreshold = 1.0
)

func init() {
	MartinegaleMovingAverageCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
	MartinegaleMovingAverageCmd.Flags().Float64VarP(&martinegaleThreshold, "threshold", "t", 1.0, "overwrite threshold value")
}

func runMartinegaleMovingAverage(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	// args
	columnName := args[0]

	// dataframe -> vector
	df := toolkit.ReadDataFrameByStdinTsv()
	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	vector.Reverse()

	for i := 1; i < len(args); i++ {
		span := toolkit.ParseInt(args[i])
		if span <= 0 {
			log.Fatalf("SPAN should be more than 1: %d\n", span)
		}

		line := vector.MartinegaleAverage(span, martinegaleThreshold)
		line.Reverse()

		newLabel := label
		if label == "" {
			newLabel = fmt.Sprintf("mma_%s_%d", columnName, span)
		}

		err = df.AddColumn(newLabel, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
