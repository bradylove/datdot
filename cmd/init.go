package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new dotter config and directory",

	Run: func(cmd *cobra.Command, args []string) {
		manager := newFileManager()
		manager.Init()

		fmt.Println("Done.")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
