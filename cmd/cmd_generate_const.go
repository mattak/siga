package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
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
	length := util.ParseInt(args[0])
	value := util.ParseFloat64(args[1])

	// header
	if label == "" {
		label = "value"
	}

	column := dataframe.CreateVectorWithValue(length, value)
	df, err := dataframe.CreateDataFrame([]string{"index", label}, []dataframe.Vector{column})
	if err != nil {
		log.Fatalf("DataFrame creation failed: %v\n", err)
	}
	df.PrintTsv(IsPreciseOutput)
}
