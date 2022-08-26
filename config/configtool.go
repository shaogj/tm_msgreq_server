package config

import (
	"github.com/mkideal/log"
)

func InitWithProviders(providers, dir string) error {
	return log.Init(providers, log.M{
		"rootdir":     dir,
		"suffix":      ".txt",
		"date_format": "%04d-%02d-%02d",
	})
}
