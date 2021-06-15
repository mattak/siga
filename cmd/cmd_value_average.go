package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	ValueAverageCmd = &cobra.Command{
		Use:     "value_average [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"va"},

		Short: "value average vector calculation",
		Long:  "value average vector calculation",
		Example: `
  siga va volume 20
  siga va volume 20 5 1
`,
		Run: runValueAverage,
	}
)

func init() {
	ValueAverageCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runValueAverage(cmd *cobra.Command, args []string) {
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

		line := vector.ValueAverage(span)
		line.Reverse()

		newLabel := label
		if label == "" {
			newLabel = fmt.Sprintf("va_%s_%d", columnName, span)
		}

		err = df.AddColumn(newLabel, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
