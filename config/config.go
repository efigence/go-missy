package config

import (
	"os"
)

// Config structure for missy
type Config struct {
	// node name. Defaults to hostname of host running the service
	NodeName string `yaml:"node_name"`
	// nrpe config path
	// if file, it will load it then try to load first include/includedir
	// if directory, it will try to load every *.cfg file in it
	// if it does not exist or it is not accessible, warning is generated
	// currently mutliple incluedirs in same file are not supported
	NrpeConfig string `yaml:"nrpe_config"`
	// Addres on which to listen for NRPE requeststs.
	// No SSL (nrpe implementation is hideously insecure and golang crypto/tls does not support that mode)
	NrpeListen string `yaml:"nrpe_listen"`
	// add self-monitoring service to sent checks
	MonitorSelf bool `yaml:"monitor_self"`
}

// New generates config with filled in defaults
func New() *Config {
	var c Config
	c.NodeName, _ = os.Hostname()
	c.MonitorSelf = true
	return &c
}
