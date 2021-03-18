package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "appengo",
	Short: "appengo command line tool",
	Long:  `Append files from the command line`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("append")
	},
}
