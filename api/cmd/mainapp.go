package cmd

import (
	"fmt"
	"squall/consts"

	"github.com/urfave/cli/v2"
)

func NewMainApp() *cli.App {
	app := cli.NewApp()
	app.Name = "squall"
	app.Usage = `squall
───────────────────────────────────────────────────────
███████╗ ██████╗ ██╗   ██╗ █████╗ ██╗     ██╗
██╔════╝██╔═══██╗██║   ██║██╔══██╗██║     ██║
███████╗██║   ██║██║   ██║███████║██║     ██║
╚════██║██║▄▄ ██║██║   ██║██╔══██║██║     ██║
███████║╚██████╔╝╚██████╔╝██║  ██║███████╗███████╗
╚══════╝ ╚══▀▀═╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝
───────────────────────────────────────────────────────`
	app.Description = `start squall server`
	app.Version = consts.AppVersion
	app.EnableBashCompletion = true
	app.DefaultCommand = CmdWeb.Name
	app.Commands = []*cli.Command{
		CmdWeb,
		CmdSeed,
	}

	return app
}

func RunMainApp(app *cli.App, args ...string) error {
	err := app.Run(args)
	if err != nil {
		return fmt.Errorf("メインコマンドの起動に失敗した: %w", err)
	}

	return nil
}
