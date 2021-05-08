package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	CountCmd = &cobra.Command{
		Use:     "count",
		Aliases: []string{"cnt"},
		Short:   "Count labels",
		Long:    "Count labels",
		Example: `
  siga count
`,
		Run: runCommandCount,
	}
)

func init() {
}

func runCommandCount(cmd *cobra.Command, args []string) {
	df := ReadDataFrameByStdinTsv()
	fmt.Println(len(df.Labels))
}
