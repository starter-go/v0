package jwt

import (
	"encoding/json"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
)

// 表示结构体(令牌)形式的 JWT
type Token struct {
	NotBefore lang.Time

	NotAfter lang.Time

	SessionID int64

	SessionUUID lang.UUID

	Properties map[string]string
}

func (inst *Token) Init() *Token {

	t := inst
	if t == nil {
		t = new(Token)
	}

	pt := t.Properties
	if pt == nil {
		pt = make(map[string]string)
		t.Properties = pt
	}

	return t
}

func (inst *Token) GetProperty(name string) string {

	if inst == nil {
		return ""
	}

	pt := inst.Properties
	if pt == nil {
		return ""
	}

	return pt[name]
}

func (inst *Token) SetProperty(name, value string) {

	if inst == nil {
		return
	}
	if name == "" || value == "" {
		return
	}

	pt := inst.Properties
	if pt == nil {
		pt = make(map[string]string)
		inst.Properties = pt
	}
}

func (inst *Token) GetTable() properties.Table {
	dst := properties.NewTable(nil)
	if inst != nil {
		src := inst.Properties
		if src != nil {
			dst.Import(src)
		}
	}
	return dst
}

func (inst *Token) String() string {
	bin, err := json.Marshal(inst)
	if err != nil {
		return "[error:" + err.Error() + "]"
	}
	return string(bin)
}
