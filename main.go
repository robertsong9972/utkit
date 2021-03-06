package main

import (
	"github.com/robertsong9972/utkit/internal/core"
	"log"

	"github.com/spf13/pflag"

	"github.com/robertsong9972/utkit/internal/config"
	"github.com/robertsong9972/utkit/internal/model"
)

type conf struct {
	confPath string
}

func main() {
	config.InitConf()
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	cfg := &conf{}
	parseFlag(cfg)
	if cfg.confPath != "" {
		config.ConfPath = cfg.confPath
		log.Println("conf_path is empty, will use default conf file in testdata/ut_package.json")
	}
	core.SysInit()
	calculator := model.NewCalculator()
	calculator.PrintCovResult()
}

func parseFlag(cfg *conf) {
	pflag.StringVar(&cfg.confPath, "conf_path", "", "json file include packages need to calculate")
	pflag.Parse()
}
