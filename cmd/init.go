package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bradylove/dotter/config"
	"github.com/onsi/say"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new dotter config and directory",

	Run: func(cmd *cobra.Command, args []string) {
		remote := say.Ask("What is the remote repository for your dot files?")

		manager := newFileManager()
		err := manager.Init(remote)
		if err != nil {
			say.Red("Init failed: %s", err)
		}

		var cfg config.Config
		viper.Unmarshal(&cfg)

		cfg.Remote = remote

		json, err := json.Marshal(&cfg)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(os.Getenv("HOME")+"/.dot/dotter", json, os.ModePerm)

		say.Green("Done.")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
