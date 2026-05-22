package filter4codec

import (
	"fmt"
	"strconv"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

type Filter4Codec struct {

	//starter:component

	_as func(subjects.FilterRegistry) //starter:as(".")

}

// Read implements subjects.ReadFilter.
func (inst *Filter4Codec) Read(c *subjects.IOC, next subjects.ReadFilterChain) error {

	// err := next.Read(c)
	// if err != nil {
	// 	return err
	// }

	// // object <<< properties
	// // inst.innerPrepareSessionDTO(c)

	// have := &c.Have
	// src := have.Properties
	// dst := have.Session
	// decoder := new(innerSessionDecoder)

	// if src == nil {
	// 	src = properties.NewTable(nil)
	// }
	// if dst == nil {
	// 	dst = new(rbac.SessionDTO)
	// }

	// err = decoder.decode(src, dst)
	// if err != nil {
	// 	return err
	// }

	// c.Have.Properties = src
	// c.Have.Session = dst
	return nil
}

// Write implements subjects.WriteFilter.
func (inst *Filter4Codec) Write(c *subjects.IOC, next subjects.WriteFilterChain) error {

	// inst.innerPrepareSessionDTO(c)

	// // object >>> properties

	// want := &c.Want
	// src := want.Session
	// dst := want.Properties
	// encoder := new(innerSessionEncoder)

	// if dst == nil {
	// 	dst = properties.NewTable(nil)
	// }
	// if src == nil {
	// 	src = new(rbac.SessionDTO)
	// }

	// err := encoder.encode(src, dst)
	// if err != nil {
	// 	return err
	// }

	// c.Want.Properties = dst
	// c.Want.Session = src

	return next.Write(c)
}

// GetRegistrationList implements subjects.FilterRegistry.
func (inst *Filter4Codec) GetRegistrationList() []*subjects.FilterRegistration {

	r1 := &subjects.FilterRegistration{
		Name:     "Filter4Codec",
		Enabled:  false,
		Priority: subjects.FilterPriorityCodec,
		Writer:   inst,
		Reader:   inst,
	}

	return []*subjects.FilterRegistration{r1}
}

// func (inst *Filter4Codec) innerPrepareSessionDTO(c *subjects.IOC) {

// 	want := &c.Want
// 	have := &c.Have

// 	if want.Session == nil {
// 		want.Session = new(rbac.SessionDTO)
// 	}

// 	if have.Session == nil {
// 		have.Session = new(rbac.SessionDTO)
// 	}

// }

func (inst *Filter4Codec) _impl() (subjects.FilterRegistry, subjects.WriteFilter, subjects.ReadFilter) {
	return inst, inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerSessionEncoder struct {
	tmp map[string]string
}

func (inst *innerSessionEncoder) encode(src *rbac.SessionDTO, dst properties.Table) error {

	if src == nil || dst == nil {
		return fmt.Errorf("param: src|dst is nil")
	}

	inst.tmp = make(map[string]string)
	inst.putAllWithSession(src)
	dst.Import(inst.tmp)
	return nil
}

func (inst *innerSessionEncoder) put(name, value string) {
	if len(name) == 0 || len(value) == 0 {
		return
	}
	inst.tmp[name] = value
}

func (inst *innerSessionEncoder) innerTryPutInt64(name subjects.PropertyName, value int64, fn func(k, v string)) {
	const base = 10
	value2 := strconv.FormatInt(value, base)
	fn(string(name), value2)
}

func (inst *innerSessionEncoder) innerTryPutString(name subjects.PropertyName, value string, fn func(k, v string)) {
	fn(string(name), value)
}

func (inst *innerSessionEncoder) innerTryPutUserID(name subjects.PropertyName, value rbac.UserID, fn func(k, v string)) {
	if value > 0 {
		inst.innerTryPutInt64(name, int64(value), fn)
	}
}

func (inst *innerSessionEncoder) innerTryPutSessionID(name subjects.PropertyName, value rbac.SessionID, fn func(k, v string)) {
	if value > 0 {
		inst.innerTryPutInt64(name, int64(value), fn)
	}
}

func (inst *innerSessionEncoder) innerTryPutTimestamp(name subjects.PropertyName, value lang.Time, fn func(k, v string)) {
	if value > 0 {
		inst.innerTryPutInt64(name, int64(value), fn)
	}
}

func (inst *innerSessionEncoder) putAllWithMap(src properties.Map) {
	if src == nil {
		return
	}
	dst := inst.tmp
	dst = src.Export(dst)
	inst.tmp = dst
}

func (inst *innerSessionEncoder) putAllWithSession(src *rbac.SessionDTO) {

	inst.innerTryPutSessionID(subjects.PNameSessionID, src.ID, inst.put)
	inst.innerTryPutUserID(subjects.PNameUserID, src.Owner, inst.put)

	inst.innerTryPutTimestamp(subjects.PNameNotBefore, src.StartedAt, inst.put)
	inst.innerTryPutTimestamp(subjects.PNameNotAfter, src.ExpiredAt, inst.put)

	inst.innerTryPutString(subjects.PNameSessionUUID, string(src.UUID), inst.put)
	inst.innerTryPutString(subjects.PNameAvatar, src.Avatar, inst.put)
	inst.innerTryPutString(subjects.PNameUserName, string(src.Username), inst.put)
	inst.innerTryPutString(subjects.PNameRoles, string(src.Roles), inst.put)
	inst.innerTryPutString(subjects.PNameEmail, string(src.Email), inst.put)
	inst.innerTryPutString(subjects.PNameLanguage, string(src.Language), inst.put)
	inst.innerTryPutString(subjects.PNameNickName, src.Nickname, inst.put)

	inst.putAllWithMap(src.Properties)
}

////////////////////////////////////////////////////////////////////////////////

type innerSessionDecoder struct {
}

func (inst *innerSessionDecoder) decode(src properties.Table, dst *rbac.SessionDTO) error {

	if src == nil || dst == nil {
		return fmt.Errorf("param: src|dst is nil")
	}

	gett := src.Getter().Optional()

	inst.innerTryGetTime(gett, subjects.PNameSessionID, func(value lang.Time) { dst.StartedAt = value })
	inst.innerTryGetTime(gett, subjects.PNameSessionID, func(value lang.Time) { dst.ExpiredAt = value })

	inst.innerTryGetInt64(gett, subjects.PNameUserID, func(value int64) { dst.Owner = rbac.UserID(value) })
	inst.innerTryGetInt64(gett, subjects.PNameSessionID, func(value int64) { dst.ID = rbac.SessionID(value) })

	inst.innerTryGetString(gett, subjects.PNameSessionUUID, func(value string) { dst.UUID = lang.UUID(value) })
	inst.innerTryGetString(gett, subjects.PNameNickName, func(value string) { dst.Nickname = value })
	inst.innerTryGetString(gett, subjects.PNameAvatar, func(value string) { dst.Avatar = value })
	inst.innerTryGetString(gett, subjects.PNameRoles, func(value string) { dst.Roles = rbac.RoleNameList(value) })

	inst.innerTryGetBool(gett, subjects.PNameRoles, func(value bool) { dst.Authenticated = value })

	return nil
}

func (inst *innerSessionDecoder) innerTryGetInt64(g properties.Getter, name subjects.PropertyName, fn func(value int64)) {
	if inst.innerContainsName(g, name) {
		value := g.GetInt64(string(name))
		fn(value)
	}
}

func (inst *innerSessionDecoder) innerTryGetString(g properties.Getter, name subjects.PropertyName, fn func(value string)) {
	value := g.GetString(string(name))
	if len(value) > 0 {
		fn(value)
	}
}

func (inst *innerSessionDecoder) innerTryGetTime(g properties.Getter, name subjects.PropertyName, fn func(value lang.Time)) {
	if inst.innerContainsName(g, name) {
		value := g.GetInt64(string(name))
		fn(lang.Time(value))
	}
}

func (inst *innerSessionDecoder) innerTryGetBool(g properties.Getter, name subjects.PropertyName, fn func(value bool)) {
	if inst.innerContainsName(g, name) {
		value := g.GetBool(string(name))
		fn(value)
	}
}

func (inst *innerSessionDecoder) innerContainsName(g properties.Getter, name subjects.PropertyName) bool {
	val := g.GetString(string(name))
	return len(val) > 0
}

////////////////////////////////////////////////////////////////////////////////

type innerPropertyMixer struct {
	table map[string]string
}

func (inst *innerPropertyMixer) init() *innerPropertyMixer {

	if inst == nil {
		inst = new(innerPropertyMixer)
	}

	if inst.table == nil {
		inst.table = make(map[string]string)
	}

	return inst
}

func (inst *innerPropertyMixer) put(name, value string) {
	if len(name) == 0 || len(value) == 0 {
		return
	}
	inst.table[name] = value
}

func (inst *innerPropertyMixer) putAllWithMap(src properties.Map) {
	if src == nil {
		return
	}
	for k, v := range src {
		inst.put(k, v)
	}
}

func (inst *innerPropertyMixer) putAllWithTable(src properties.Table) {
	if src == nil {
		return
	}
	keys := src.Names()
	for _, name := range keys {
		value := src.GetProperty(name)
		inst.put(name, value)
	}
}

func (inst *innerPropertyMixer) result() properties.Table {
	pt := properties.NewTable(nil)
	pt.Import(inst.table)
	return pt
}

////////////////////////////////////////////////////////////////////////////////
// EOF
