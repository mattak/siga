package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	AnomalyDetectionCmd = &cobra.Command{
		Use:     "anomaly_detection [COLUMN_NAME] [SPAN] [SIGMA_THRESHOLD]",
		Aliases: []string{"ad"},

		Short: "Anomaly detection",
		Long:  "Anomaly detection",
		Example: `
anomaly detection by 1.5 sigma, 10 data points:
  siga ad volume 10 1.5 < sample.tsv
`,
		Run: runCommandAnomalyDetection,
	}
)

func init() {
}

func runCommandAnomalyDetection(cmd *cobra.Command, args []string) {
	if len(args) < 3 {
		log.Fatal("COLUMN_NAME, SPAN, THRESHOLD should be declared")
	}

	columnName := args[0]
	span := ParseInt(args[1])
	sigmaThreshold := ParseFloat64(args[2])
	df := ReadDataFrameByStdinTsv()

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatalf("Not found index of header name: %s\n", columnName)
	}
	vector.Reverse()

	result := vector.SigmaAnomalies(span, sigmaThreshold)
	result.Reverse()
	err = df.AddColumn("anomalies", result)
	if err != nil {
		log.Fatalf("add result column failed\n")
	}

	df.PrintTsv()
}
