package root

import (
	"fmt"
	"os"

	"github.com/ecsctl/ecsctl/internal/commands/config"
	"github.com/ecsctl/ecsctl/internal/commands/get"
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ecsctl",
	Short: "ecsctl controls an AWS ECS cluster.",
	DisableFlagsInUseLine: true,
	Version: "v0.0.1",
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add commands to root command
	rootCmd.AddCommand(get.NewCmdGet())
	rootCmd.AddCommand(config.NewCmdConfig())
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(home+"/.ecsctl/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	// Load config into the configutils package
	configutils.LoadConfig()
}