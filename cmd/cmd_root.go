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
	label = ""
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
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(IdentityCmd)
	RootCmd.AddCommand(AnomalyCmd)
	RootCmd.AddCommand(MeansCmd)
	RootCmd.AddCommand(DeviationsCmd)
	RootCmd.AddCommand(MultiplyCmd)
	RootCmd.AddCommand(DivideCmd)
	RootCmd.AddCommand(SelectColumnCmd)
	RootCmd.AddCommand(SelectRecordCmd)
	RootCmd.AddCommand(ReverseCmd)
	RootCmd.AddCommand(InvertCmd)
	RootCmd.AddCommand(TakeCmd)
	RootCmd.AddCommand(ConstCmd)
	RootCmd.AddCommand(ShiftCmd)
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
	RootCmd.AddCommand(SumCmd)
	RootCmd.AddCommand(SubCmd)
	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(RenameCmd)
	RootCmd.AddCommand(ProfitFactorCmd)
	RootCmd.AddCommand(PayoffRatioCmd)
	RootCmd.AddCommand(TradingEvaluationCmd)
	RootCmd.AddCommand(DollarCostAverageCmd)
	RootCmd.AddCommand(GenerateCmd)
}
