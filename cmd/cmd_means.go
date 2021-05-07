package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	MeansCmd = &cobra.Command{
		Use:     "means [COLUMN_NAME] [POINTS]+",
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
	if len(args) < 3 {
		log.Fatal("COLUMN_NAME, BASE_LINE, TARGET_LINE should be declared")
	}

	// args
	columnName := args[0]

	// dataframe -> vector
	df := ReadDataFrameByStdinTsv()
	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatal(err)
	}
	vector.Reverse()

	for i := 1; i < len(args); i++ {
		points := ParseInt(args[i])
		if points <= 0 {
			log.Fatalf("POINTS should be more than 1: %d\n", points)
		}

		line := vector.Means(points)
		line.Reverse()
		err = df.AddColumn(fmt.Sprintf("mean_%s_%d", columnName, points), line)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv()
}
