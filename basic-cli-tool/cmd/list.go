/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/glebarez/sqlite"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tableData := pterm.TableData{
			{"First Name","Last Name","Age"},
		}
		db, _ := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db",Dbname)), &gorm.Config{})
		users := []User{}
		db.Model(User{}).Find(&users)
		for _,user := range users{
			tableData = append(tableData, []string{user.FirstName,user.LastName,strconv.Itoa(user.Age)})
		}
		pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(tableData).Render()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
