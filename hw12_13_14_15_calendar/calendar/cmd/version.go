package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of application",
	Long:  `All software has versions. This is application`,
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersion()
	},
}
