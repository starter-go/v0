package boot

import (
	"sort"

	"github.com/starter-go/application"
	"github.com/starter-go/v0/libvlog/api/config"
	"github.com/starter-go/vlog"
)

type VLogBootLoader struct {

	//starter:component

	_as func(application.Lifecycle) //starter:as(".")

	FilterList    []vlog.FilterRegistry //starter:inject(".")
	ConfigService config.Service        //starter:inject("#")
}

// Create implements vlog.LoggerFactory.
func (inst *VLogBootLoader) Create() vlog.Logger {

	fcloader := new(innerFilterChainLoader)
	fcloader.init(inst.FilterList)
	chain := fcloader.load()
	cfgser := inst.ConfigService

	cfg, _ := cfgser.GetConfiguration()
	if cfg == nil {
		cfg = new(config.Configuration)
		cfgser.LoadDefault(cfg)
	}

	ada := new(vlog.LoggerAdapter)
	ada.SetTargetChain(chain)
	ada.SetSender(inst)
	ada.SetLevelAccepted(cfg.Level)
	ada.SetTag(cfg.Tag)

	return ada
}

// Life implements application.Lifecycle.
func (inst *VLogBootLoader) Life() *application.Life {
	return &application.Life{
		Order:      -9999,
		OnStartPre: inst.startup,
	}
}

func (inst *VLogBootLoader) startup() error {
	vlog.SetLoggerFactory(inst)
	return nil
}

func (inst *VLogBootLoader) _impl() (application.Lifecycle, vlog.LoggerFactory) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerFilterChainLoader struct {
	items []*vlog.FilterRegistration
}

func (inst *innerFilterChainLoader) init(list []vlog.FilterRegistry) {

	for _, it1 := range list {
		if it1 == nil {
			continue
		}
		tmp := it1.ListLogFilterRegistration()
		for _, it2 := range tmp {
			inst.addItem(it2)
		}
	}

}

func (inst *innerFilterChainLoader) addItem(it *vlog.FilterRegistration) {

	if it == nil {
		return
	}

	if it.Filter == nil {
		return
	}

	inst.items = append(inst.items, it)
}

func (inst *innerFilterChainLoader) Len() int {
	return len(inst.items)
}
func (inst *innerFilterChainLoader) Less(i1, i2 int) bool {
	all := inst.items
	o1 := all[i1].Order
	o2 := all[i2].Order
	return (o1 < o2)
}
func (inst *innerFilterChainLoader) Swap(i1, i2 int) {
	all := inst.items
	all[i1], all[i2] = all[i2], all[i1]
}

func (inst *innerFilterChainLoader) sort() {
	sort.Sort(inst)
}

func (inst *innerFilterChainLoader) load() vlog.FilterChain {

	inst.sort()

	all := inst.items
	cb := new(vlog.GroupFilterChainBuilder)
	cb.AddRegistration(all...)

	// for _, it := range all {
	// 	cb.AddFilter(it.Filter)
	// }

	return cb.Build(vlog.GroupMain)
}

////////////////////////////////////////////////////////////////////////////////
