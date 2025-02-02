package day5

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
    day := &cobra.Command{
    Use: "2024-5",
        Short: "Problems for Day 5",
    }

    day.AddCommand(aCmd)
    day.AddCommand(bCmd)

    root.AddCommand(day)
}
