package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	MeansCmd = &cobra.Command{
		Use:     "means [COLUMN_NAME] [SPAN]+",
		Aliases: []string{"mn"},

		Short: "means vector calculation",
		Long:  "means vector calculation",
		Example: `
  siga mn volume 20
  siga mn volume 20 5 1
`,
		Run: runCommandMeans,
	}
)

func init() {
}

func runCommandMeans(cmd *cobra.Command, args []string) {
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

		line := vector.Means(span)
		line.Reverse()
		err = df.AddColumn(fmt.Sprintf("mean_%s_%d", columnName, span), line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
