package cmd

import (
	"github.com/mattak/siga/pkg"
	"github.com/spf13/cobra"
	"log"
)

var (
	GenerateRandomCmd = &cobra.Command{
		Use:     "rand [LENGTH] [FROM] [TO]",
		Aliases: []string{"r"},

		Short: "Generate random sequence",
		Long:  "Generate random sequence",
		Example: `
generate range (0,1) values
  siga gen rand 10 0 1
`,
		Run: runCommandGenerateRandom,
	}
)

func init() {
	GenerateRandomCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandGenerateRandom(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		log.Fatal("[LENGTH] [FROM] [TO] should be declared")
	}
	length := pkg.ParseInt(args[0])
	fromValue := pkg.ParseFloat64(args[1])
	toValue := pkg.ParseFloat64(args[2])

	// header
	if label == "" {
		label = "value"
	}

	column := pkg.CreateVectorWithRandom(length, fromValue, toValue)
	df := pkg.CreateDataFrame([]string{"index", label}, []pkg.Vector{column})
	df.PrintTsv(IsPreciseOutput)
}
