/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tubskns/icos-shell/client/pkg/openapi"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "icos-shell",
	Short: "icos-shell - a CLI to interface the ICOS Shell",
	Long: `icos-shell - a CLI to interface the ICOS Shell
   
The icos-shell can be used to modify or inspect resources in the ICOS controller from the terminal`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing icos-shell '%s'", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$XDG_CONFIG_HOME/icos-shell/config.yaml", "config file")
	rootCmd.PersistentFlags().StringP("lighthouse", "l", "", "lighthouse address")
	rootCmd.PersistentFlags().StringP("controller", "b", "", "controller address")
	rootCmd.PersistentFlags().StringP("username", "u", "", "username parameter")
	rootCmd.PersistentFlags().StringP("password", "p", "", "password parameter")
	rootCmd.PersistentFlags().StringP("token", "t", "", "token")

	viper.BindPFlag("lighthouse", rootCmd.PersistentFlags().Lookup("lighthouse"))
	viper.BindPFlag("controller", rootCmd.PersistentFlags().Lookup("controller"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Config:", viper.ConfigFileUsed())
	}

	if viper.GetString("controller") == "" {
		fmt.Println("Controller not defined")
		os.Exit(0)
	} else if viper.GetString("controller") != "" && viper.GetString("token") == "" {
		fmt.Println("Token is empty, you must login first")
	} else {
		fmt.Println("Controller:", viper.GetString("controller"))
		fmt.Println("Token found:", viper.GetString("token"))
	}
	openapi.Init(viper.GetString("controller"))
}