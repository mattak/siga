package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	SelectRecordCmd = &cobra.Command{
		Use:     "select_record [LABEL_NAME|INDEX]+",
		Aliases: []string{"sr"},

		Short: "select record",
		Long:  "Select record",
		Example: `
  siga sr 0 3 5 < sample.tsv
  siga sr 2020-05-01 < sample.tsv
`,
		Run: runCommandSelectRecord,
	}
)

func init() {
}

func runCommandSelectRecord(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("LABEL_NAME or NUMBER should be declared")
	}

	df := dataframe.ReadDataFrameByStdinTsv()
	indexes := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		index := df.FindLabelIndex(args[i])
		if index != -1 {
			indexes[i] = index
		} else {
			index = util.ParseInt(args[i])
			if index >= 0 {
				indexes[i] = index
			} else {
				newIndex := len(df.Labels) + index
				if newIndex < 0 {
					log.Fatalf("Index is not found: %d is mapped %d\n", index, newIndex)
				}
				indexes[i] = newIndex
			}
		}
	}
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
