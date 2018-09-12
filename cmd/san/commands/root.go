package commands

import (
	"github.com/spf13/cobra"
)

var debug bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Display debug information")
}

// RootCmd is the san CLI's root command.
var RootCmd = &cobra.Command{
	Use:   "san",
	Short: "the Simple And Needed configuration format",
	Long:  "the Simple And Needed configuration format. https://astrocorp.net/san",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
