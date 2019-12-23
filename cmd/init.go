package cmd

import (
	"github.com/alzedd/golb/internal/pkg/settings"
	"github.com/spf13/cobra"
)

var s *settings.Settings

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
