/*
Package cmd Liciense
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// location code name
// A 10 臺北市
// B 11 臺中市
// C 12 基隆市
// D 13 臺南市
// E 14 高雄市
// F 15 新北市
// G 16 宜蘭縣
// H 17 桃園市
// I 34 嘉義市
// J 18 新竹縣
// K 19 苗栗縣
// M 21 南投縣
// N 22 彰化縣
// O 35 新竹市
// P 23 雲林縣
// Q 24 嘉義縣
// T 27 屏東縣
// U 28 花蓮縣
// V 29 臺東縣
// W 32 金門縣
// X 30 澎湖縣
// Z 33 連江縣
// L 20 臺中市
// R 25 臺南市
// S 26 高雄市
// Y 31
var locationCode map[byte]int = map[byte]int{
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
	'G': 16,
	'H': 17,
	'I': 34,
	'J': 18,
	'K': 19,
	'M': 21,
	'N': 22,
	'O': 35,
	'P': 23,
	'Q': 24,
	'T': 27,
	'U': 28,
	'V': 29,
	'W': 32,
	'X': 30,
	'Z': 33,
	'L': 20,
	'R': 25,
	'S': 26,
	'Y': 31,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "twidgen",
	Short: "Creating fake Taiwan ID",
	Long:  `twidgen is a CLI tool for creating fake Taiwan ID`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		location := byte('A' + rand.Int31n(25))
		sex := rand.Intn(2) + 1
		serials := make([]int, 7)
		for i := 0; i < 7; i++ {
			serials[i] = rand.Intn(10)
		}

		locationNum := locationCode[location]
		sum := (locationNum / 10) + ((locationNum % 10) * 9) + sex*8
		for i := 0; i < 7; i++ {
			sum += serials[i] * (7 - i)
		}
		lastNum := (10 - (sum % 10)) % 10

		id := string(location) + fmt.Sprintf("%v", sex)
		for i := 0; i < 7; i++ {
			id += fmt.Sprintf("%v", serials[i])
		}
		id += fmt.Sprintf("%v", lastNum)

		fmt.Println(id)
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.twidgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".twidgen" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".twidgen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
