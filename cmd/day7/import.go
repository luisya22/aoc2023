package day7

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
    day := &cobra.Command{
        Use: "7",
        Short: "Problems for Day 7",
    }

    day.AddCommand(aCmd)
    day.AddCommand(bCmd)

    root.AddCommand(day)
}