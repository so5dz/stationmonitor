package main

import (
	"log"

	"github.com/so5dz/aprsis"
	smconfig "github.com/so5dz/stationmonitor/config"
	"github.com/so5dz/utils/config"
	"github.com/so5dz/utils/misc"
)

func main() {
	misc.WrapMain(mainWithError)()
}

func mainWithError() error {
	cfg, err := config.LoadConfigFromArgs[smconfig.StationMonitorConfig]()
	if err != nil {
		return err
	}

	is := aprsis.APRSIS{}
	is.Start(cfg.APRSIS.Host, cfg.APRSIS.Port, cfg.APRSIS.User, cfg.APRSIS.Pass, cfg.APRSIS.Filter)
	is.OnPacket(func(tnc2Bytes []byte) {
		log.Println(string(tnc2Bytes))
	})

	log.Println("Started StationMonitor", cfg)
	misc.BlockUntilInterrupted()
	return nil
}
