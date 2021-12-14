package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var inputFileName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "geocode",
	Short: "Tool to get geocode given a csv file",
	Long:  `geocode will go through each record in the csv and call geocode api and return a result file with census record`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFileName, _ = cmd.Flags().GetString("file")
		if inputFileName == "" {
			inputFileName = "input"
		}
		ExecuteTask()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("file", "f", "", "Specify the input file")
	_ = rootCmd.MarkPersistentFlagRequired("file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".geocode" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".geocode")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
