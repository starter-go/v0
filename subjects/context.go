package subjects

import (
	"context"

	"github.com/starter-go/application/attributes"
)

type Context struct {
	CC context.Context

	Cache       *Cache // for reader
	CacheHolder CacheHolder
	CacheLoader CacheLoader

	Buffer *Buffer // for writer

	FlagCreate bool
	FlagDirty  bool

	Facade Subject

	Checker Checker

	Reader ReadFilterChain

	Writer WriteFilterChain

	ChainHolder FilterChainHolder

	Attributes attributes.Table
}

func (inst *Context) Clone() *Context {

	c2 := new(Context)
	*c2 = *inst
	c2.Facade = NewSubject(c2)
	return c2
}

func (inst *Context) GetBuffer(autoMake bool) *Buffer {
	b := inst.Buffer
	if b == nil && autoMake {
		b = new(Buffer)
		inst.Buffer = b
	}
	return b
}

func (inst *Context) GetCC(autoMake bool) context.Context {
	cc := inst.CC
	if cc == nil && autoMake {
		cc = context.Background()
		inst.CC = cc
	}
	return cc
}

func (inst *Context) GetAttributes(autoMake bool) attributes.Table {
	att := inst.Attributes
	if att == nil && autoMake {
		att = attributes.NewTable(nil)
		inst.Attributes = att
	}
	return att
}

func (inst *Context) GetCache(autoMake bool) *Cache {
	c := inst.Cache
	if c == nil && autoMake {
		c = new(Cache)
		inst.Cache = c
	}
	return c
}

func (inst *Context) GetCacheHolder(autoMake bool) CacheHolder {
	ch := inst.CacheHolder
	if ch == nil && autoMake {
		ch = new(innerCacheHolderImpl)
		inst.CacheHolder = ch
	}
	return ch
}

func (inst *Context) GetChecker(autoMake bool) Checker {
	ckr := inst.Checker
	if ckr == nil && autoMake {
		checker, err := makeNewChecker(inst)
		if err != nil {
			panic(err)
		}
		ckr = checker
		inst.Checker = checker
	}
	return ckr
}

func (inst *Context) GetCacheLoader(autoMake bool) CacheLoader {
	cl := inst.CacheLoader
	if cl == nil && autoMake {
		cl = new(innerCacheLoaderImpl)
		inst.CacheLoader = cl
	}
	return cl
}

func (inst *Context) GetReader(autoMake bool) ReadFilterChain {
	ch := inst.Reader
	if ch == nil && autoMake {
		rdr, err := inst.ChainHolder.GetReaderChain()
		if err != nil {
			panic(err)
		}
		inst.Reader = rdr
		ch = rdr
	}
	return ch
}

func (inst *Context) GetWriter(autoMake bool) WriteFilterChain {
	ch := inst.Writer
	if ch == nil && autoMake {
		wtr, err := inst.ChainHolder.GetWriterChain()
		if err != nil {
			panic(err)
		}
		inst.Writer = wtr
		ch = wtr
	}
	return ch
}

func (inst *Context) GetFacade(autoMake bool) Subject {
	res := inst.Facade
	if res == nil && autoMake {
		facade := &innerSubjectFacade{
			context: inst,
		}
		res = facade
		inst.Facade = facade
	}
	return res
}

func (inst *Context) NewIOC(method Method) *IOC {

	c2 := new(IOC)
	buffer := inst.Buffer

	c2.CC = inst.CC
	c2.Context = inst
	c2.Want.Method = method

	if buffer != nil {
		c2.Want.Properties = buffer.Properties
	}

	return c2
}
