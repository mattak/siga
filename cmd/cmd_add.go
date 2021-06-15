package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var (
	AddCmd = &cobra.Command{
		Use:     "add [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]+",
		Aliases: []string{},
		Short:   "Add",
		Long:    "Add",
		Example: `
  siga add column1 column2
  siga add column1 5
`,
		Run: runCommandAdd,
	}
)

func init() {
	AddCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandAdd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than 1 COLUMN_NAME should be declared")
	}
	df := pkg.ReadDataFrameByStdinTsv()

	matrix := make(pkg.Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = pkg.CreateVector(len(df.Labels))

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

	vector := matrix.Add()

	if label == "" {
		label = fmt.Sprintf("add_%s", strings.Join(args, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	df.PrintTsv(IsPreciseOutput)
}
