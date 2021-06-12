package cmd

import (
	"github.com/mattak/siga/toolkit"
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
	length := toolkit.ParseInt(args[0])
	startValue := toolkit.ParseFloat64(args[1])
	rwalks := make([]toolkit.RwalkValue, len(args)-2)

	for i := 2; i < len(args); i++ {
		index := (i - 2)
		rwalks[index] = toolkit.CreateRwalkValue(args[i], ":")
	}

	rwalkSetting := toolkit.CreateRwalkSetting(startValue, rwalks)

	// header
	if label == "" {
		label = "value"
	}

	column := toolkit.CreateVectorWithRandomWalk(length, rwalkSetting)
	df := toolkit.CreateDataFrame([]string{"index", label}, []toolkit.Vector{column})
	df.PrintTsv(IsPreciseOutput)
}
