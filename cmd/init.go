package cmd

import (
	"github.com/onsi/say"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new datdot config and directory",

	Run: func(cmd *cobra.Command, args []string) {
		remote := say.Ask("What is the remote repository for your dot files?")

		manager := newFileManager()
		err := manager.Init(remote)
		if err != nil {
			say.Red("Init failed: %s", err)
		}

		say.Green("Done.")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
