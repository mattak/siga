package cmd

import (
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
)

var (
	IdentityCmd = &cobra.Command{
		Use:     "identity",
		Aliases: []string{"id"},

		Short: "Identity mapping",
		Long:  "Identity mapping",
		Example: `
mapping from stdin:
  siga id < sample.tsv

sample data format:
  date	open	close	high	low	volume
  2020-01-01	100	105	110	90	100
  2020-01-02	100	105	110	90	100
`,
		Run: runCommandIdentity,
	}
)

func init() {
}

func runCommandIdentity(cmd *cobra.Command, args []string) {
	df := toolkit.ReadDataFrameByStdinTsv()
	df.PrintTsv(IsPreciseOutput)
}
