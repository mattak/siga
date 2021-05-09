package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	VersionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{},

		Short: "show version",
		Long:  "show version",
		Example: `
  siga version
`,
		Run: runCommandVersion,
	}
)

func init() {
}

func runCommandVersion(cmd *cobra.Command, args []string) {
	fmt.Println(VERSION)
}