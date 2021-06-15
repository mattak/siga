package cmd

import (
	"github.com/mattak/siga/pkg"
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
	length := pkg.ParseInt(args[0])
	startValue := pkg.ParseFloat64(args[1])
	rwalks := make([]pkg.RwalkValue, len(args)-2)

	for i := 2; i < len(args); i++ {
		index := (i - 2)
		rwalks[index] = pkg.CreateRwalkValue(args[i], ":")
	}

	rwalkSetting := pkg.CreateRwalkSetting(startValue, rwalks)

	// header
	if label == "" {
		label = "value"
	}

	column := pkg.CreateVectorWithRandomWalk(length, rwalkSetting)
	df := pkg.CreateDataFrame([]string{"index", label}, []pkg.Vector{column})
	df.PrintTsv(IsPreciseOutput)
}
