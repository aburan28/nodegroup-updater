package main

import (
	"nodegroup-updater/cmd"

	"github.com/spf13/viper"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	regions := viper.GetString("regions")
	clusters := viper.GetString("clusters")

}
