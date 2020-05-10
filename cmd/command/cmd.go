package command

import (
	"apimake/cmd/action"
	"apimake/cmd/flag"
	"apimake/pkg/constant"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Init() {
	app := new(cli.App)
	app.Name = constant.AppName
	app.Usage = constant.AppUsage
	app.Description = constant.AppUsage
	app.Flags = append(app.Flags, flag.VersionFlag, flag.LanguageFlag, flag.FileFlag)
	app.Commands = append(app.Commands, createApiCommand, updateApiCommand, listApiCommand, showApiCommand)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	createApiCommand = &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "create api base info",
		Action:  action.CreateApi,
	}

	updateApiCommand = &cli.Command{
		Name:        "update",
		Aliases:     []string{"edit"},
		Description: "update api info, such as: header, request, response",
		Usage:       "update api info, such as: header, request, response",
		Flags: []cli.Flag{
			flag.UpdateTypeFlag,
			flag.ApiIdFlag,
		},
		Action: action.UpdateApi,
	}

	listApiCommand = &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "show total api items list",
		Action:  action.ListApi,
	}

	showApiCommand = &cli.Command{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "show api detail info",

		Action: action.ShowApi,
		Flags: []cli.Flag{
			flag.ApiIdFlag,
		},
	}
)
