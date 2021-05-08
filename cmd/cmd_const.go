package cmd

import (
	"fmt"
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
}

func runCommandConst(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	for i := 0; i < len(args); i++ {
		n := ParseFloat64(args[i])
		vector := CreateVector(len(df.Labels))
		vector.Fill(n)
		err := df.AddColumn(fmt.Sprintf("const_%.3f", n), vector)
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv()
}
