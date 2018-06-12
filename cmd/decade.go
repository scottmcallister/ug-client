package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decadeCmd)
}

var decadeCmd = &cobra.Command{
	Use:   "decade",
	Short: "Search for tabs by decade",
	Long:  "Search for tabs by decade",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Search for tabs by decade")
	},
}
