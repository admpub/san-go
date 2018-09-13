package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astroflow/astroflow-go/log"
	"github.com/spf13/cobra"
)

var fmtOutput string

func init() {
	RootCmd.AddCommand(FmtCmd)
	FmtCmd.Flags().StringVarP(&fmtOutput, "output", "o", "", "Place the output into <file>")
}

// FmtCmd is the `fmt` command. It permit to automatically format a SAN file
var FmtCmd = &cobra.Command{
	Use:   "fmt [file]",
	Args:  cobra.ExactArgs(1),
	Short: "automatically formats a SAN file",
	Long:  "automatically formats a SAN file",
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

		data = format(data)

		if fmtOutput == "" {
			fmtOutput = file
		}
		err = ioutil.WriteFile(fmtOutput, data, fileInfo.Mode())
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func format(data []byte) []byte {
	buf := bytes.NewBuffer(data)
	var ret bytes.Buffer
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		trimedLine := strings.TrimSpace(scanner.Text())
		fmt.Println(trimedLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}

	ret.Write(data)
	return ret.Bytes()
}
