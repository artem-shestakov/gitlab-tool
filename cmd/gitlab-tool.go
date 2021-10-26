package main

import (
	"log"
	"os"

	"github.com/artem-shestakov/gitlab-tool.git/internal/pipeline"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Commands = []*cli.Command{
		&pipeline.Pipeline,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
