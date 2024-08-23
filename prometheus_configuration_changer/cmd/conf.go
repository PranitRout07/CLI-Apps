/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(InputConfig("demo","/metrics",[]string{"localhost:9090"}))
		job_name := pterm.DefaultInteractiveTextInput
		job_name.DefaultText = "Enter job name" 
		jn,_ := job_name.Show()

		metrics_path := pterm.DefaultInteractiveTextInput
		metrics_path.DefaultText = "Enter metrics path" 
		mp,_ := metrics_path.Show()
	
		targets := pterm.DefaultInteractiveTextInput
		targets.DefaultText = "Enter the url in a list" 
		tg,_ := targets.Show()


		Output := InputConfig(jn,mp,tg)
		pterm.Info.Print(Output)
		
		res := exec.Command("docker","exec","-u","root","prometheus","sh","-c",fmt.Sprintf(`echo "%s" >> /etc/prometheus/prometheus.yml`,Output))
		out,_ := res.StdoutPipe()
		res.Start()
		output,_ := io.ReadAll(out)
		pterm.Info.Print(string(output))
		res.Wait()



		res1 := exec.Command("docker","restart","prometheus")
		out1,_ := res1.StdoutPipe()
		res1.Start()
		output1,_:= io.ReadAll(out1)
		pterm.Info.Print("Prometheus restarted..",string(output1))
		res1.Wait()
	},
}

func init() {
	rootCmd.AddCommand(confCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func InputConfig(job_name string,metrics_path string,targets string)string {

	return fmt.Sprintf(`  - job_name: "\"%s\""

    metrics_path: '%s'


    static_configs:
      - targets: ["\"%v\""]`,job_name,metrics_path,targets)
	



  }