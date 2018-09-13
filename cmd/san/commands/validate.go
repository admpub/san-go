package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astroflow/astroflow-go/log"
	"github.com/phasersec/san-go/parser"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ValidateCmd)
}

// ValidateCmd is the `validate` command. It permit to check if a .san file is valid or not
var ValidateCmd = &cobra.Command{
	Use:   "valdiate [file]",
	Args:  cobra.ExactArgs(1),
	Short: "check if a .san file validity",
	Long:  "check if a .san file validity",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fileInfo, err := os.Stat(file)
		if err != nil {
			log.Fatal(fmt.Sprintf("error: opening %s: %s", file, err))
		}

		if fileInfo.IsDir() {
			log.Fatal("error: " + file + " is a directory")
		}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = parser.Parse(data)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}
