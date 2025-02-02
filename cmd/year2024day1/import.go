package day1

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
    day := &cobra.Command{
    Use: "2024-1",
        Short: "Problems for Day 1",
    }

    day.AddCommand(aCmd)
    day.AddCommand(bCmd)

    root.AddCommand(day)
}
