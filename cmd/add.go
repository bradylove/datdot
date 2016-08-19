package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

		fmt.Println("add called")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
