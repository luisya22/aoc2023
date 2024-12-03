package day2

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
    day := &cobra.Command{
    Use: "2024-2",
        Short: "Problems for Day 2",
    }

    day.AddCommand(aCmd)
    day.AddCommand(bCmd)

    root.AddCommand(day)
}
