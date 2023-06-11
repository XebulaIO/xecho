package xecho

import "github.com/XebulaIO/gommon/log"

var logger = log.New("-")

func init() {
	logger.SetSkip(3)
}
