package pipeline

import "github.com/spf13/cobra"

type CobraCommandInput struct {
	Cmd  *cobra.Command
	Args []string
}
