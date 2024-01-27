package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"squall/config"
	"squall/fetcher"
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
	go func() {
		err = r.Run(config.Config.Address)
		if err != nil {
			log.Fatal("routerの起動に失敗した: %w", err)
		}
	}()

	b, err := os.ReadFile("./feeds.yml")
	if err != nil {
		return fmt.Errorf("フィードリストの読み込みに失敗した: %w", err)
	}
	go func() {
		err = fetcher.Run(bytes.NewReader(b))
		if err != nil {
			log.Fatal("フィード取得に失敗した: %w", err)
		}
	}()

	select {}

	return nil
}
