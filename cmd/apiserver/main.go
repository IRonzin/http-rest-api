package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/IRonzin/http-rest-api/internal/app/apiserver"
)

var (
	configPath  string
	isNeedPprof bool
)

func init() {
	flag.StringVar(&configPath, "config-path", "C:/Users/VladimirRonzin/source/repos/http-rest-api/configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if "true" == os.Getenv("IS_NEED_PPROF") {
		isNeedPprof = true
	}

	if err := apiserver.Start(config, isNeedPprof); err != nil {
		log.Fatal(err)
	}

}
