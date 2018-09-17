package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astroflow/astroflow-go/log"
	"github.com/phasersec/san-go"
	"github.com/phasersec/san-go/lexer"
	"github.com/spf13/cobra"
)

var fmtOutput string

func init() {
	RootCmd.AddCommand(FmtCmd)
	FmtCmd.Flags().StringVarP(&fmtOutput, "output", "o", "", "Place the output into <file>")
}

// FmtCmd is the `fmt` command. It permit to automatically format a SAN file
// TODO:
var FmtCmd = &cobra.Command{
	Use:   "fmt [file]",
	Args:  cobra.ExactArgs(1),
	Short: "automatically formats a SAN file",
	Long:  "automatically formats a SAN file",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fileInfo, err := os.Stat(file)
		s := map[string]interface{}{}
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

		err = san.Unmarshal(data, &s)
		if err != nil {
			log.Fatal(err.Error())
		}

		data, err = san.Marshal(s)
		if err != nil {
			log.Fatal(err.Error())
		}

		if fmtOutput == "" {
			fmtOutput = file
		}
		err = ioutil.WriteFile(fmtOutput, data, fileInfo.Mode())
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

// TODO:
func format(data []byte) []byte {

	var ret bytes.Buffer
	lx := lexer.NewLexer(data)

	for token := range lx.Next() {
		ret.WriteString(token.Value)
		ret.WriteRune(' ')
		eq := <-lx.Next()
		ret.WriteString(eq.Value)
		ret.WriteRune(' ')
	}

	return ret.Bytes()
}
