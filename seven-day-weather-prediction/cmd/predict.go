/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

type PredictResponse struct {
	City struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Coord struct {
			Lon float64 `json:"lon,omitempty"`
			Lat float64 `json:"lat,omitempty"`
		} `json:"coord,omitempty"`
		Country    string `json:"country,omitempty"`
		Population int    `json:"population,omitempty"`
		Timezone   int    `json:"timezone,omitempty"`
	} `json:"city,omitempty"`
	Cod     string  `json:"cod,omitempty"`
	Message float64 `json:"message,omitempty"`
	Cnt     int     `json:"cnt,omitempty"`
	List    []struct {
		Dt      int `json:"dt,omitempty"`
		Sunrise int `json:"sunrise,omitempty"`
		Sunset  int `json:"sunset,omitempty"`
		Temp    struct {
			Day   float64 `json:"day,omitempty"`
			Min   float64 `json:"min,omitempty"`
			Max   float64 `json:"max,omitempty"`
			Night float64 `json:"night,omitempty"`
			Eve   float64 `json:"eve,omitempty"`
			Morn  float64 `json:"morn,omitempty"`
		} `json:"temp,omitempty"`
		FeelsLike struct {
			Day   float64 `json:"day,omitempty"`
			Night int     `json:"night,omitempty"`
			Eve   float64 `json:"eve,omitempty"`
			Morn  float64 `json:"morn,omitempty"`
		} `json:"feels_like,omitempty"`
		Pressure int `json:"pressure,omitempty"`
		Humidity int `json:"humidity,omitempty"`
		Weather  []struct {
			ID          int    `json:"id,omitempty"`
			Main        string `json:"main,omitempty"`
			Description string `json:"description,omitempty"`
			Icon        string `json:"icon,omitempty"`
		} `json:"weather,omitempty"`
		Speed  float64 `json:"speed,omitempty"`
		Deg    int     `json:"deg,omitempty"`
		Gust   float64 `json:"gust,omitempty"`
		Clouds int     `json:"clouds,omitempty"`
		Pop    int     `json:"pop,omitempty"`
	} `json:"list,omitempty"`
}

// predictCmd represents the predict command
var predictCmd = &cobra.Command{
	Use:   "predict",
	Short: "Predict weather for a particular city",
	Long:  `Predict weather for a particular city`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("predict called")
		city, err := cmd.Flags().GetString("city")
		if err != nil {
			pterm.Info.Print("Error while getting city name '--city' flag")
		}
		API_KEY, val := os.LookupEnv("WEATHER_API_KEY")
		if !val {
			pterm.Info.Print("Error while reading the API key.")
		}

		resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast/daily?q=%s&cnt=3&appid=%s", city, API_KEY))
		pterm.Info.Print(fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast/daily?q=%s&cnt=3&appid=%s", city, API_KEY))
		if err != nil {
			pterm.Info.Print("Error while getting response")
			return
		}
		defer resp.Body.Close()
		// var v interface{}
		// json.Unmarshal(resp.Body,v)
		pterm.Info.Print("Body response:", resp.StatusCode)
	},
}

func init() {
	rootCmd.AddCommand(predictCmd)
	predictCmd.PersistentFlags().String("city", "", "A city name.")
}
