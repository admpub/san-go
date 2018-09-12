package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/astroflow/astroflow-go/log"
	"github.com/go-yaml/yaml"
	"github.com/phasersec/san-go"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ConvertCmd)
}

// ConvertCmd is the `convert` command. It permit to convert configuration file in other formats to SAN
var ConvertCmd = &cobra.Command{
	Use:   "convert [file]",
	Args:  cobra.ExactArgs(1),
	Short: "convert a .(toml|yml|json) file to a .san one",
	Long:  "convert a .(toml|yml|json) file to a .san one",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fileInfo, err := os.Stat(file)
		if err != nil {
			log.Fatal(fmt.Sprintf("error: opening %s: %s", file, err))
		}

		if fileInfo.IsDir() {
			log.Fatal("error: " + file + " is a directory")
		}

		ext := filepath.Ext(file)
		fileWithotuExt := strings.TrimSuffix(file, ext)
		data, err := ioutil.ReadFile(file)

		s := map[string]interface{}{}
		switch ext {
		case ".toml":
			err = toml.Unmarshal(data, &s)
		case ".yml", ".yaml":
			err = yaml.Unmarshal(data, &s)
		case ".json":
			err = json.Unmarshal(data, &s)
		default:
			log.Fatal("extension " + ext + " not recgnized, valid extension are one of the following: [toml, yml, yaml, json]")
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		dataToWrite, err := san.Marshal(s)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = ioutil.WriteFile(fileWithotuExt+".san", dataToWrite, fileInfo.Mode())
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}
