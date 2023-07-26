package cmd

import "github.com/spf13/cobra"

func AddSysCommands(root *cobra.Command) {
	root.AddCommand(initAppCommand())
	root.AddCommand(initCronCommand())
	root.AddCommand(initEnvCommand())
}
