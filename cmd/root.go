package cmd

// This code generated by go generate.
// DO NOT EDIT BY HAND!

import (
    "os"

    "github.com/spf13/cobra"
    
    "github.com/luisya22/aoc2023/cmd/day1"
    "github.com/luisya22/aoc2023/cmd/day2"
    "github.com/luisya22/aoc2023/cmd/day3"
    "github.com/luisya22/aoc2023/cmd/day4"
)

func addDays(root *cobra.Command){
    
    day1.AddCommandsTo(root)
    day2.AddCommandsTo(root)
    day3.AddCommandsTo(root)
    day4.AddCommandsTo(root)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "Advent of Code 2023 Solutions",
	Long:  `Golang implementations for the 2023 Advent of Code problems.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aoc2023.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	addDays(rootCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
