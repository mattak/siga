package cmd

import (
	"github.com/spf13/cobra"
)

var (
	GenerateCmd = &cobra.Command{
		Use:     "generate [SUB_COMMAND]",
		Aliases: []string{"gen"},

		Short: "Generate sequence",
		Long:  "Generate sequence",
		Example: `
  siga gen const 10 2 < sample.tsv
`,
	}
)

func init() {
	GenerateCmd.AddCommand(GenerateConstCmd)
}
