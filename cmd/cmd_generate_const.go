package cmd

import (
	"github.com/mattak/siga/pkg"
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
	length := pkg.ParseInt(args[0])
	value := pkg.ParseFloat64(args[1])

	// header
	if label == "" {
		label = "value"
	}

	column := pkg.CreateVectorWithValue(length, value)
	df := pkg.CreateDataFrame([]string{"index", label}, []pkg.Vector{column})
	df.PrintTsv(IsPreciseOutput)
}
