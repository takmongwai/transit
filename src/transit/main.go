package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[DEBUG] ", 1)
}
func main() {
	cfg, err := LoadConfigFile(SelfRoot() + "/etc/dev_config.json")
	if err != nil {
		logger.Println(err)
	}
	logger.Println(cfg)
}
