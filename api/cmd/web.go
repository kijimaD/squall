package cmd

import (
	"fmt"
	"squall/config"
	"squall/models"
	"squall/routers"

	"github.com/urfave/cli/v2"
)

var CmdWeb = &cli.Command{
	Name:   "web",
	Usage:  "start server",
	Action: runWeb,
	Flags:  []cli.Flag{},
}

func runWeb(_ *cli.Context) error {
	models.GetDB()
	r, err := routers.NewRouter()
	if err != nil {
		return fmt.Errorf("routerの起動に失敗した: %w", err)
	}
	err = r.Run(config.Config.Address)
	if err != nil {
		return fmt.Errorf("routerの起動に失敗した: %w", err)
	}

	return nil
}
