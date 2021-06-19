package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
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
	length := util.ParseInt(args[0])
	fromValue := util.ParseFloat64(args[1])
	toValue := util.ParseFloat64(args[2])

	// header
	if label == "" {
		label = "value"
	}

	column := dataframe.CreateVectorWithRandom(length, fromValue, toValue)
	df, err := dataframe.CreateDataFrame([]string{"index", label}, []dataframe.Vector{column})
	if err != nil {
		log.Fatalf("DataFrame creation failed: %v\n", err)
	}
	df.PrintTsv(IsPreciseOutput)
}
