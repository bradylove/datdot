package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "dotter",
	Short: "dotter is a simple dotfile manager",
	Long: `dotter is a simple dotfile manager that uses a git repo for syncing
your dotfiles.`,
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("dotter")     // name of config file (without extension)
	viper.AddConfigPath("$HOME/.dot") // adding home directory as first search path

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
