package cmd

import (
	"github.com/mattak/siga/toolkit"
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
  siga gen const 10 2
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
	length := toolkit.ParseInt(args[0])
	value := toolkit.ParseFloat64(args[1])

	// header
	if label == "" {
		label = "value"
	}

	column := toolkit.CreateVectorWithValue(length, value)
	df := toolkit.CreateDataFrame([]string{"index", label}, []toolkit.Vector{column})
	df.PrintTsv(IsPreciseOutput)
}
