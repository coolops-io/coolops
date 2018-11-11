package main

import (
	"fmt"
	"os"

	"github.com/coolopsio/coolops/command"
	"github.com/coolopsio/coolops/flags"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:      "build:new",
		Usage:     "Notify coolops.io about a new build",
		ArgsUsage: "[build name]",
		Flags: []cli.Flag{
			cli.GenericFlag{
				Name:  "metadata, m",
				Usage: "Information to be sent as field on the Slack message (`name=value`)",
				Value: &flags.KeyValueFlag{
					Values: make(map[string]string),
				},
			},
			cli.GenericFlag{
				Name:  "param, p",
				Usage: "Build parameters to be injected in the deployment container (`name=value`)",
				Value: &flags.KeyValueFlag{
					Values: make(map[string]string),
				},
			},
			cli.StringFlag{
				Name:   "token, t",
				Usage:  "The project's api token",
				EnvVar: "COOLOPS_PROJECT_API_TOKEN",
			},
		},
		Action: command.CmdNewBuild,
	},
	{
		Name:  "build:new:circleci",
		Usage: "Notify coolops.io about a new build using default conventions for CircleCI builds",
		Flags: []cli.Flag{
			cli.GenericFlag{
				Name:  "metadata, m",
				Usage: "Information to be sent as field on the Slack message (`name=value`)",
				Value: &flags.KeyValueFlag{
					Values: make(map[string]string),
				},
			},
			cli.GenericFlag{
				Name:  "param, p",
				Usage: "Build parameters to be injected in the deployment container (`name=value`)",
				Value: &flags.KeyValueFlag{
					Values: make(map[string]string),
				},
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "The build name. By default it's [branch-name]-[build-number]",
			},
			cli.StringFlag{
				Name:   "token, t",
				Usage:  "The project's api token",
				EnvVar: "COOLOPS_PROJECT_API_TOKEN",
			},
		},
		Action: command.CmdNewBuildCircleCI,
	},
	{
		Name:  "build:new:gitlab",
		Usage: "Notify coolops.io about a new build using default conventions for GitLab builds",
		Flags: []cli.Flag{
			cli.GenericFlag{
				Name:  "metadata, m",
				Usage: "Information to be sent as field on the Slack message (`name=value`)",
				Value: &flags.KeyValueFlag{
					Values: make(map[string]string),
				},
			},
			cli.GenericFlag{
				Name:  "param, p",
				Usage: "Build parameters to be injected in the deployment container (`name=value`)",
				Value: &flags.KeyValueFlag{
					Values: make(map[string]string),
				},
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "The build name. By default it's [branch-name]-[build-number]",
			},
			cli.StringFlag{
				Name:   "token, t",
				Usage:  "The project's api token",
				EnvVar: "COOLOPS_PROJECT_API_TOKEN",
			},
		},
		Action: command.CmdNewBuildGitlab,
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
