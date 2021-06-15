package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	HarmonicMovingAverageCmd = &cobra.Command{
		Use:     "harmonic_moving_average [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"hma"},

		Short: "harmonic moving average vector calculation",
		Long:  "harmonic moving average vector calculation",
		Example: `
  siga hma volume 20
  siga hma volume 20 5 1
`,
		Run: runHarmonicMovingAverage,
	}
)

func init() {
	HarmonicMovingAverageCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runHarmonicMovingAverage(cmd *cobra.Command, args []string) {
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

		line := vector.HarmonicMovingAverage(span)
		line.Reverse()

		newLabel := label
		if label == "" {
			newLabel = fmt.Sprintf("hma_%s_%d", columnName, span)
		}

		err = df.AddColumn(newLabel, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
