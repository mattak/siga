package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "siga",
		Short: "Signal Analyzer for sequential data",
		Long:  "Signal Analyzer for sequential data",
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(IdentityCmd)
	RootCmd.AddCommand(AnomalyCmd)
	RootCmd.AddCommand(MeansCmd)
	RootCmd.AddCommand(DeviationsCmd)
	RootCmd.AddCommand(MultiplyCmd)
	RootCmd.AddCommand(SelectCmd)
	RootCmd.AddCommand(ReverseCmd)
	RootCmd.AddCommand(TakeCmd)
	RootCmd.AddCommand(ConstCmd)
	RootCmd.AddCommand(EqualCmd)
	RootCmd.AddCommand(FilterCmd)
	RootCmd.AddCommand(FilterNotCmd)
	RootCmd.AddCommand(GreaterEqualCmd)
	RootCmd.AddCommand(GreaterThanCmd)
	RootCmd.AddCommand(DetectAnyCmd)
	RootCmd.AddCommand(DetectAllCmd)
	RootCmd.AddCommand(CountCmd)
}
