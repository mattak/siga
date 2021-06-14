package cmd

import (
	"github.com/spf13/cobra"
)

var (
	FilterCmd = &cobra.Command{
		Use:     "filter [SUB_COMMAND]",
		Aliases: []string{"fl"},
		Short:   "Filter sequence by value",
		Long:    "Filter sequence by value",
	}
)

func init() {
	FilterCmd.AddCommand(FilterOrCmd)
	FilterCmd.AddCommand(FilterAndCmd)
	FilterCmd.AddCommand(FilterNotAndCmd)
	FilterCmd.AddCommand(FilterNotOrCmd)
}
