package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	AnomalyCmd = &cobra.Command{
		Use:     "anomaly [COLUMN_NAME] [SPAN]",
		Aliases: []string{"an"},

		Short: "Anomaly calculation",
		Long:  "Anomaly calculation. output is relative to sigma `(v-mean)/σ`. 1.0 means 1σ",
		Example: `
anomaly calculation by 10 data points:
  siga ad volume 10 < sample.tsv
`,
		Run: runCommandAnomaly,
	}
)

func init() {
	AnomalyCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandAnomaly(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnName := args[0]
	span := ParseInt(args[1])
	df := ReadDataFrameByStdinTsv()

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatalf("Not found index of header name: %s\n", columnName)
	}
	vector.Reverse()

	result := vector.SigmaAnomalies(span)
	result.Reverse()

	if label == "" {
		label = fmt.Sprintf("%s_anomaly", columnName)
	}
	err = df.AddColumn(label, result)
	if err != nil {
		log.Fatalf("add result column failed\n")
	}

	df.PrintTsv(IsPreciseOutput)
}
