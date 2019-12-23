package cmd

import (
	"fmt"
	"os"

	"github.com/alzedd/golb/internal/pkg/settings"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golb",
	Short: "",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var s *settings.Settings
	cobra.OnInitialize(s.ReadConfig)
}
