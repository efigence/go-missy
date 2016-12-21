package main

import (
	"github.com/XANi/go-yamlcfg"
	"github.com/efigence/go-missy/config"
	"github.com/op/go-logging"
	"os"
)

var version string
var log = logging.MustGetLogger("main")
var stdout_log_format = logging.MustStringFormatter("%{color:bold}%{time:2006-01-02T15:04:05.9999Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")

var cfgFiles = []string{
	"$HOME/.config/missy/config.yaml",
	//
	"/etc/missy/config.yaml",
	// for local tests
	"./cfg/config.yaml",
	"./cfg/config.default.yaml",
}

func main() {
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	stderrFormatter := logging.NewBackendFormatter(stderrBackend, stdout_log_format)
	logging.SetBackend(stderrFormatter)
	logging.SetFormatter(stdout_log_format)
	log.Infof("Starting missy version: %s", version)
	cfg := config.New()

	err := yamlcfg.LoadConfig(cfgFiles, cfg)
	if err != nil {
		log.Panicf("Error loading config: %s", err)
	}

}
