package cmd

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
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
	pipe := pipeline.CobraCommandInput{cmd, args}.CreateCountPipeInt()
	df := dataframe.ReadDataFrameByStdinTsv()
	value := pipe.Execute(df)
	fmt.Println(value)
}
