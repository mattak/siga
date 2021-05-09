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
	IsPreciseOutput = false
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(
		&IsPreciseOutput,
		"precise", "p",
		false,
		"floating number should be print precisely",
	)
	RootCmd.AddCommand(IdentityCmd)
	RootCmd.AddCommand(AnomalyCmd)
	RootCmd.AddCommand(MeansCmd)
	RootCmd.AddCommand(DeviationsCmd)
	RootCmd.AddCommand(MultiplyCmd)
	RootCmd.AddCommand(SelectCmd)
	RootCmd.AddCommand(ReverseCmd)
	RootCmd.AddCommand(TakeCmd)
	RootCmd.AddCommand(ConstCmd)
	RootCmd.AddCommand(FilterCmd)
	RootCmd.AddCommand(FilterNotCmd)
	RootCmd.AddCommand(EqualCmd)
	RootCmd.AddCommand(GreaterEqualCmd)
	RootCmd.AddCommand(GreaterThanCmd)
	RootCmd.AddCommand(LessEqualCmd)
	RootCmd.AddCommand(LessThanCmd)
	RootCmd.AddCommand(DetectAnyCmd)
	RootCmd.AddCommand(DetectAllCmd)
	RootCmd.AddCommand(CountCmd)
}
