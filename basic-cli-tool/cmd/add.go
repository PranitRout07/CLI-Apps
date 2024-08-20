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

// structure
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int
}

const Dbname = "test"
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		
		firstname := pterm.DefaultInteractiveTextInput
		firstname.DefaultText = "Enter your firstname: "
		fn, _ := firstname.Show()

		lastname := pterm.DefaultInteractiveTextInput
		lastname.DefaultText = "Enter your lastname: "
		ln, _ := lastname.Show()

		age := pterm.DefaultInteractiveTextInput
		age.DefaultText = "Enter your age: "
		a, _ := age.Show()

		db, _ := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", Dbname)), &gorm.Config{})
		db.AutoMigrate(&User{})

		db.Create(&User{
			FirstName: fn, LastName: ln, Age: func() int { x, _ := strconv.Atoi(a); return x }(),
		})

		pterm.Info.Print("Saved Info To Database")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
