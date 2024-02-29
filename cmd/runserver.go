package cmd

import (
	"github.com/atpuxiner/grapi/cmd/internal/runserver"
	"github.com/spf13/cobra"
)

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "runserver",
	Long:  `runserver`,
	Run: func(cmd *cobra.Command, args []string) {
		runserver.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)
	// add flags
	runserverCmd.Flags().StringP("port", "p", "", "端口")
}
