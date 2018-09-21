package main

import (
	"github.com/astrocorp42/astroflow-go"
	"github.com/astrocorp42/astroflow-go/log"
	"github.com/phasersec/san-go/cmd/san/commands"
)

func main() {
	log.Config(
		astroflow.SetFormatter(astroflow.NewCLIFormatter()),
		astroflow.SetLevel(astroflow.InfoLevel),
	)

	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
