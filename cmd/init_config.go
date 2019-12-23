package cmd

import (
	"github.com/alzedd/golb/internal/pkg/settings"
	"github.com/spf13/cobra"
)

var initconfigCmd = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		settings.WriteConfig(true)
	},
}

func init() {
	initCmd.AddCommand(initconfigCmd)
}
