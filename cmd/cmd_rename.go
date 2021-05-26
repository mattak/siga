package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	RenameCmd = &cobra.Command{
		Use:     "rename_column ([OLD_COLUMN_NAME] [NEW_COLUMN_NAME])+",
		Aliases: []string{"rc"},

		Short: "Rename column",
		Long:  "Rename column",
		Example: `
  siga rc column_old column_new < sample.tsv
`,
		Run: runCommandRename,
	}
)

func init() {
}

func runCommandRename(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("More than two OLD_NAME, NEW_NAME should be declared")
	}
	if len(args)%2 != 0 {
		log.Fatal("Argument length should be even")
	}

	df := ReadDataFrameByStdinTsv()
	headers := df.Headers

	for i := 0; i < len(args); i += 2 {
		oldName := args[i]
		newName := args[i+1]
		for j := 0; j < len(headers); j++ {
			if headers[j] == oldName {
				headers[j] = newName
			}
		}
	}
	df.Headers = headers

	df.PrintTsv(IsPreciseOutput)
}
