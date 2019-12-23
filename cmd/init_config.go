package cmd

import (
	"github.com/spf13/cobra"
)

var initconfigCmd = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		s.WriteConfig(true)
	},
}

func init() {
	initCmd.AddCommand(initconfigCmd)
}
