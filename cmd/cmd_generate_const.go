package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	GenerateConstCmd = &cobra.Command{
		Use:     "const [LENGTH] [VALUE]",
		Aliases: []string{"c"},

		Short: "Generate const sequence",
		Long:  "Generate const sequence",
		Example: `
  siga gen const 10 2 < sample.tsv
`,
		Run: runCommandGenerateConst,
	}
)

func init() {
	GenerateConstCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandGenerateConst(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("[LENGTH] [VALUE] should be declared")
	}
	length := ParseInt(args[0])
	value := ParseFloat64(args[1])

	// header
	if label == "" {
		label = "value"
	}

	column := CreateVectorWithValue(length, value)
	df := CreateDataFrame([]string{"index", label}, []Vector{column})
	df.PrintTsv(IsPreciseOutput)
}
