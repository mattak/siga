package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	DeviationsCmd = &cobra.Command{
		Use:     "deviations [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"dv"},

		Short: "Deviation",
		Long:  "Deviation",
		Example: `
deviation by span 5
  siga dv close 5 < sample.tsv
deviation by span 5, 10
  siga dv close 5 10 < sample.tsv
`,
		Run: runCommandDeviations,
	}
)

func init() {
}

func runCommandDeviations(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnName := args[0]

	df := ReadDataFrameByStdinTsv()
	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	vector.Reverse()

	for i := 1; i < len(args); i++ {
		span := ParseInt(args[i])
		if span <= 0 {
			log.Fatalf("SPAN should be more than 1: %d\n", span)
		}

		line := vector.Deviations(span)
		line.Reverse()
		err = df.AddColumn(fmt.Sprintf("deviation_%s_%d", columnName, span), line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
