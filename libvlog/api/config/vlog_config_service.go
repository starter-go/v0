package config

import "github.com/starter-go/vlog"

type Configuration struct {
	Level vlog.Level
	Tag   string
}

type Service interface {
	GetConfiguration() (*Configuration, error)

	LoadDefault(cfg *Configuration)
}
