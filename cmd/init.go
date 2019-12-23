package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
