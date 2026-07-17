package iconfig

import (
	"github.com/starter-go/v0/libvlog/api/config"
	"github.com/starter-go/vlog"
)

type VLogConfigService struct {

	//starter:component

	_as func(config.Service) //starter:as("#")

	StrLevel string //starter:inject("${vlog.level}")
	StrTag   string //starter:inject("${vlog.tag}")

	cache *config.Configuration
}

// LoadDefault implements [config.Service].
func (inst *VLogConfigService) LoadDefault(cfg *config.Configuration) {

	if cfg == nil {
		return
	}

	cfg.Level = vlog.INFO
	cfg.Tag = "libvlog"
}

// GetConfiguration implements [config.Service].
func (inst *VLogConfigService) GetConfiguration() (*config.Configuration, error) {
	c := inst.cache
	if c == nil {
		c2, err := inst.innerLoadConfiguration()
		if err != nil {
			return nil, err
		}
		c = c2
		inst.cache = c2
	}
	return c, nil
}

func (inst *VLogConfigService) innerLoadConfiguration() (*config.Configuration, error) {

	level, err := vlog.ParseLevel(inst.StrLevel)
	if err != nil {
		return nil, err
	}

	cfg := new(config.Configuration)
	cfg.Level = level
	cfg.Tag = inst.StrTag
	return cfg, nil
}

func (inst *VLogConfigService) _impl() config.Service {
	return inst
}
