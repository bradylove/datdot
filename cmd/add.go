package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <filepaths>",
	Short: "add a file to your dotfiles",

	Long: `add will move a file to your dotter directory, create a symlink in its
original location, and commit that file to the local git repository. Use sync to
push changes to your dot files to the remote repository`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("requires filepath(s)")
			os.Exit(1)
		}

		// TODO: Ensure dotter has been initialized

		manager := newFileManager()

		if err := manager.Add(args[0]); err != nil {
			fmt.Printf("Failed to add file to your dotfiles: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Done.")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
