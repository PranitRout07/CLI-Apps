/*
Copyright © 2024 NAME HERE <pranitrout72@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use:   "weather-app-cli",
	Short: "Predict Weather",
	Long: `Predict weather upto 16 days ahead by giving only a city name as input.`,

}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


