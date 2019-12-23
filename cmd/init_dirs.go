package cmd

import (
	"github.com/alzedd/golb/internal/pkg/fsutils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var initdirsCmd = &cobra.Command{
	Use:   "dirs",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		fsutils.MkDirs(fs, s)
	},
}

func init() {
	initCmd.AddCommand(initdirsCmd)
}
