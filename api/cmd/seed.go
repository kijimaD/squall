package cmd

import (
	"log"
	"squall/factories"
	"squall/models"

	"github.com/urfave/cli/v2"
)

var CmdSeed = &cli.Command{
	Name:   "seed",
	Usage:  "start server",
	Action: runSeed,
	Flags:  []cli.Flag{},
}

func runSeed(_ *cli.Context) error {
	models.GetDB()

	entries := []models.Entry{}
	count := models.GetDB().Find(&entries).RowsAffected
	if count == 1 {
		log.Println("seed実行済みのため、スキップ")

		return nil
	}

	var deps []factories.Dependency
	_, deps = factories.MakeEntry(factories.Fields{"URL": "https://amazon.com", "Title": "amazon"}, deps)
	_, deps = factories.MakeEntry(factories.Fields{"URL": "https://google.com", "Title": "google"}, deps)
	for _, m := range deps {
		err := models.GetDB().Create(m).Error
		if err != nil {
			return err
		}
	}

	return nil
}
