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
	label           = ""
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

	// dataframe
	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(AnomalyCmd)
	RootCmd.AddCommand(ConstCmd)
	RootCmd.AddCommand(DeviationsCmd)
	RootCmd.AddCommand(DivideCmd)
	RootCmd.AddCommand(EqualCmd)
	RootCmd.AddCommand(FilterCmd)
	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(GreaterEqualCmd)
	RootCmd.AddCommand(GreaterThanCmd)
	RootCmd.AddCommand(HarmonicMovingAverageCmd)
	RootCmd.AddCommand(IdentityCmd)
	RootCmd.AddCommand(InvertCmd)
	RootCmd.AddCommand(LessEqualCmd)
	RootCmd.AddCommand(LessThanCmd)
	RootCmd.AddCommand(MartinegaleMovingAverageCmd)
	RootCmd.AddCommand(MultiplyCmd)
	RootCmd.AddCommand(NormalizeCmd)
	RootCmd.AddCommand(RenameCmd)
	RootCmd.AddCommand(ReverseCmd)
	RootCmd.AddCommand(ReverseTakeCmd)
	RootCmd.AddCommand(SelectColumnCmd)
	RootCmd.AddCommand(SelectLastValueCmd)
	RootCmd.AddCommand(SelectRecordCmd)
	RootCmd.AddCommand(ShiftCmd)
	RootCmd.AddCommand(SimpleMovingAverageCmd)
	RootCmd.AddCommand(SubCmd)
	RootCmd.AddCommand(SumCmd)
	RootCmd.AddCommand(TakeCmd)
	RootCmd.AddCommand(ValueAverageCmd)

	// int
	RootCmd.AddCommand(CountCmd)

	// bool
	RootCmd.AddCommand(DetectAnyCmd)
	RootCmd.AddCommand(DetectAllCmd)

	// etc
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(ProfitFactorCmd)
	RootCmd.AddCommand(PayoffRatioCmd)
	RootCmd.AddCommand(TradingEvaluationCmd)
	RootCmd.AddCommand(DollarCostAverageCmd)
}
