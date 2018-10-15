package commands

import (
	"fmt"

	"github.com/bloom42/san-go"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(VersionCmd)
}

// VersionCmd display the version information
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "display version information",
	Long:  "display version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", san.Version)
	},
}
