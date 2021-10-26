package pipeline

import "github.com/urfave/cli/v2"

var (
	Pipeline = cli.Command{
		Name:  "pipeline",
		Usage: "Pipeline tools",
		Subcommands: []*cli.Command{
			{
				Name:  "notify",
				Usage: "Send notification about pipeline status",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "channel",
						Usage: "Channel for sending notification",
					},
				},
				Action: notify,
			},
		},
	}
)
