package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "siga",
		Short: "Signal analyzer for sequential data",
		Long:  "Signal analyzer for sequential data",
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(IdentityMapCmd)
	RootCmd.AddCommand(AnomalyDetectionCmd)
	RootCmd.AddCommand(MeansCmd)
	RootCmd.AddCommand(SelectCmd)
	RootCmd.AddCommand(DeviationsCmd)
	RootCmd.AddCommand(MultiplyCmd)
}
