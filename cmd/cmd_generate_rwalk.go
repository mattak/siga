package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	GenerateRwalkCmd = &cobra.Command{
		Use:     "rwalk [LENGTH] [START_VALUE] ([WEIGHT],[FROM],[TO])+",
		Aliases: []string{"rw"},

		Short: "Generate random walk sequence",
		Long:  "Generate random walk sequence",
		Example: `
generate random walk values
  siga gen rwalk 10 100 0.5:0:0.1 0.5:-0.1:0
  siga gen rwalk 10 100 0.15:0.05:0.2 0.15:-0.2:-0.05 0.7:-0.05:0.05
`,
		Run: runCommandGenerateRandomWalk,
	}
)

func init() {
	GenerateRwalkCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandGenerateRandomWalk(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("[LENGTH] [START_VALUE] ([WEIGHT],[FROM],[TO]) should be declared")
	}
	length := util.ParseInt(args[0])
	startValue := util.ParseFloat64(args[1])
	rwalks := make([]dataframe.RwalkValue, len(args)-2)

	for i := 2; i < len(args); i++ {
		index := (i - 2)
		rwalks[index] = dataframe.CreateRwalkValue(args[i], ":")
	}

	rwalkSetting := dataframe.CreateRwalkSetting(startValue, rwalks)

	// header
	if label == "" {
		label = "value"
	}

	column := dataframe.CreateVectorWithRandomWalk(length, rwalkSetting)
	df, err := dataframe.CreateDataFrame([]string{"index", label}, []dataframe.Vector{column})
	if err != nil {
		log.Fatalf("DataFrame creation failed: %v\n", err)
	}
	df.PrintTsv(IsPreciseOutput)
}
