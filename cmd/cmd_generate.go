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
	}
)

func init() {
	GenerateCmd.AddCommand(GenerateConstCmd)
	GenerateCmd.AddCommand(GenerateRandomCmd)
}
