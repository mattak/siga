package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	GoldenCrossCmd = &cobra.Command{
		Use:     "golden_cross [COLUMN_NAME] [BASE_LINE] [TARGET_LINE]",
		Aliases: []string{"gc"},

		Short: "Golden cross detection",
		Long:  "Golden cross detection",
		Example: `
5point mean line cross against 20point mean line:
  siga gc volume 20 5
`,
		Run: runCommandGoldenCross,
	}
	RetryLimit = 10
)

func init() {
}

func runCommandGoldenCross(cmd *cobra.Command, args []string) {
	if len(args) < 3 {
		log.Fatal("COLUMN_NAME, BASE_LINE, TARGET_LINE should be declared")
	}

	// args
	columnName := args[0]
	baseLinePoints := ParseInt(args[1])
	targetLinePoints := ParseInt(args[2])

	if baseLinePoints <= 0 {
		log.Fatal("BASE_LINE should be more than 1")
	}
	if targetLinePoints <= 0 {
		log.Fatal("TARGET_LINE should be more than 1")
	}

	// dataframe
	df := ReadDataFrameByStdinTsv()

	if len(df.Labels)-1 < baseLinePoints {
		log.Fatal("BASE_LINE should be more than 1")
	}
	if len(df.Labels)-1 < targetLinePoints {
		log.Fatal("TARGET_LINE should be more than 1")
	}

	// open,close,high,low,volume
	columnIndex := df.FindHeaderIndex(columnName) - 1
	if columnIndex < 0 {
		log.Fatalf("Not found index of header name: %s\n", columnName)
	}
	length := len(df.Labels)

	// calc
	baseLineMean1 := df.Mean(columnIndex, length-baseLinePoints-1, baseLinePoints)
	baseLineMean2 := df.Mean(columnIndex, length-baseLinePoints, baseLinePoints)
	targetLineMean1 := df.Mean(columnIndex, length-targetLinePoints-1, targetLinePoints)
	targetLineMean2 := df.Mean(columnIndex, length-targetLinePoints, targetLinePoints)

	cross1 := baseLineMean1 < targetLineMean1
	cross2 := baseLineMean2 < targetLineMean2
	isCrossing := cross1 != cross2

	fmt.Printf(
		"crossing:%t\tup1:%t\tup2:%t\tbase1:%.2f\tbase2:%.2f\ttarget1:%.2f\ttarget2:%.2f\n",
		isCrossing,
		cross1,
		cross2,
		baseLineMean1,
		baseLineMean2,
		targetLineMean1,
		targetLineMean2,
	)
}
