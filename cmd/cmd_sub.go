package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var (
	SubCmd = &cobra.Command{
		Use:     "sub [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]+",
		Aliases: []string{},
		Short:   "Subtract",
		Long:    "Subtract",
		Example: `
  siga sub column1 column2
  siga sub column1 5
`,
		Run: runCommandSub,
	}
)

func init() {
	SubCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandSub(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than 1 COLUMN_NAME should be declared")
	}
	df := dataframe.ReadDataFrameByStdinTsv()

	matrix := make(dataframe.Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = dataframe.CreateVector(len(df.Labels))

		vector, err := df.ExtractColumn(args[i])
		if err == nil {
			matrix[i] = vector
		} else {
			v, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
			} else {
				matrix[i].Fill(v)
			}
		}
	}

	vector := matrix.Subtract()

	if label == "" {
		label = fmt.Sprintf("sub_%s", strings.Join(args, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
