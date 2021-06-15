package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	ConstCmd = &cobra.Command{
		Use:     "const [NUMBER]+",
		Aliases: []string{"c"},

		Short: "Const",
		Long:  "Const",
		Example: `
add const 1
  siga c 1 < sample.tsv
`,
		Run: runCommandConst,
	}
)

func init() {
	ConstCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandConst(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("NUMBER should be declared")
	}

	df := dataframe.ReadDataFrameByStdinTsv()
	for i := 0; i < len(args); i++ {
		n := util.ParseFloat64(args[i])
		vector := dataframe.CreateVector(len(df.Labels))
		vector.Fill(n)

		if label == "" {
			label = fmt.Sprintf("const_%.3f", n)
		}

		err := df.AddColumn(label, vector)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
