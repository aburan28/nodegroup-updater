package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "nodegroup-updater",
		Short: "Run the nodegroup-updater script",
		Long:  `This script is used to update the nodegroups of an EKS cluster`,
	}
	cfgFile string
)

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nodegroup-updater")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("regions", "us-east-1", "The AWS regions to look for clusters in")
	viper.BindPFlag("regions", rootCmd.PersistentFlags().Lookup("regions"))
	rootCmd.MarkPersistentFlagRequired("regions")
	rootCmd.PersistentFlags().String("clusters", "", "The EKS clusters to update")
	viper.BindPFlag("clusters", rootCmd.PersistentFlags().Lookup("clusters"))
	rootCmd.MarkPersistentFlagRequired("clusters")
}
