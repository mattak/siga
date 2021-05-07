package cmd

import (
	"github.com/spf13/cobra"
)

var (
	IdentityMapCmd = &cobra.Command{
		Use:     "identity_map",
		Aliases: []string{"im"},

		Short: "Identity mapping",
		Long:  "Identity mapping",
		Example: `
mapping from stdin:
  siga im < sample.tsv

sample data format:
  date	open	close	high	low	volume
  2020-01-01	100	105	110	90	100
  2020-01-02	100	105	110	90	100
`,
		Run: runCommandIdentityMap,
	}
)

func init() {
}

func runCommandIdentityMap(cmd *cobra.Command, args []string) {
	df := ReadDataFrameByStdinTsv()
	df.PrintTsv()
}
