/*
Copyright © 2024 NAME HERE <pranitrout72@gmail.com>
*/
package cmd

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		
		querry := args[0]
		fmt.Println(querry)
		cm := exec.Command("awk",fmt.Sprintf("/%s/",querry),"./cmd/data.json")
		//filename
		//Querry
		res, _ := cm.StdoutPipe()

		cm.Start()
		out, _ := io.ReadAll(res)
		fmt.Println(string(out))
		cm.Wait()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().String("search","","Search")
	// searchCmd.Flags().StringP("search", "s", "", "Help message for search")
}
