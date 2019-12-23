package cmd

import (
	"github.com/alzedd/golb/internal/pkg/fsutils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var initallCmd = &cobra.Command{
	Use:   "all",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		s.WriteConfig(true)
		fsutils.MkDirs(fs, s)
	},
}

func init() {
	initCmd.AddCommand(initallCmd)
}
