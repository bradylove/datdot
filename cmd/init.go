package cmd

import (
	"fmt"
	"github.com/bradylove/dotter/filemanager"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new dotter config and directory",

	Run: func(cmd *cobra.Command, args []string) {
		manager := filemanager.New(os.Getenv("HOME"))
		manager.Init()

		fmt.Println("Done.")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
