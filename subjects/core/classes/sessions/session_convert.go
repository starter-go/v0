package sessions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/localization"
	"github.com/starter-go/v0/subjects"
)

// DTO <===> Entity

func ConvertD2E(src *DTO, dst *Entity) error {

	tmp := properties.NewTable(nil)
	err := ConvertD2P(src, tmp)
	if err != nil {
		return err
	}
	return ConvertP2E(tmp, dst)

}

func ConvertE2D(src *Entity, dst *DTO) error {

	tmp := properties.NewTable(nil)
	err := ConvertE2P(src, tmp)
	if err != nil {
		return err
	}
	return ConvertP2D(tmp, dst)

}

func ConvertListE2D(src []*Entity, dst []*DTO) ([]*DTO, error) {

	for _, o1 := range src {
		o2 := new(DTO)
		err := ConvertE2D(o1, o2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}

	return dst, nil
}

// Entity <===> Properties

func ConvertP2E(src properties.Table, dst *Entity) error {

	t0a := lang.Time(0)
	t0b := t0a.Time()
	dst.NotBefore = t0b
	dst.NotAfter = t0b

	getter := new(innerPTGetter)
	getter.init(src)

	getter.Name(subjects.PNameSessionID).GetID(func(value int64) { dst.ID = rbac.SessionID(value) })
	getter.Name(subjects.PNameSessionUUID).GetString(func(value string) { dst.UUID = lang.UUID(value) })
	getter.Name(subjects.PNameUserID).GetID(func(value int64) { dst.Owner = rbac.UserID(value) })
	getter.Name(subjects.PNameRoles).GetString(func(value string) { dst.Roles = rbac.RoleNameList(value) })
	getter.Name(subjects.PNameCreatedAt).GetTime(func(value lang.Time) { dst.CreatedAt = value.Time() })
	getter.Name(subjects.PNameUpdatedAt).GetTime(func(value lang.Time) { dst.UpdatedAt = value.Time() })
	getter.Name(subjects.PNameNotBefore).GetTime(func(value lang.Time) { dst.NotBefore = value.Time() })
	getter.Name(subjects.PNameNotAfter).GetTime(func(value lang.Time) { dst.NotAfter = value.Time() })
	getter.Name(subjects.PNameMaxTokenAge).GetInt64(func(value int64) { dst.MaxTokenAge = lang.Milliseconds(value) })
	getter.Name(subjects.PNameAuthenticated).GetBool(func(value bool) { dst.Authenticated = value })

	getter.GetExt(func(value map[string]string) {
		pm := properties.Map(value)
		dst.Properties = pm.Text()
	})

	return nil
}

func ConvertE2P(src *Entity, dst properties.Table) error {

	setter := new(innerPTSetter)
	setter.init(dst)
	setter.PutMap(src.Properties.Map())

	setter.Name(subjects.PNameSessionID).PutID(int64(src.ID))
	setter.Name(subjects.PNameSessionUUID).PutString(src.UUID.String())
	setter.Name(subjects.PNameUserID).PutID(int64(src.Owner))

	setter.Name(subjects.PNameCreatedAt).PutTime2t(src.CreatedAt)
	setter.Name(subjects.PNameUpdatedAt).PutTime2t(src.UpdatedAt)
	setter.Name(subjects.PNameNotBefore).PutTime2t(src.NotBefore)
	setter.Name(subjects.PNameNotAfter).PutTime2t(src.NotAfter)
	setter.Name(subjects.PNameMaxTokenAge).PutTimeMS(src.MaxTokenAge)

	setter.Name(subjects.PNameRoles).PutString(string(src.Roles))
	setter.Name(subjects.PNameAuthenticated).PutBool(src.Authenticated)

	// setter.Name(subjects.PNameUserName).PutString(src.Username.String())
	// setter.Name(subjects.PNameAvatar).PutString(src.Avatar)
	// setter.Name(subjects.PNameEmail).PutString(src.Email.String())
	// setter.Name(subjects.PNameNickName).PutString(src.Nickname)
	// setter.Name(subjects.PNameLanguage).PutString(src.Language.String())
	// setter.Name(subjects.PNameDeletedAt).PutTime2t(src.DeletedAt)

	return nil
}

// DTO <===> Properties

func ConvertP2D(src properties.Table, dst *DTO) error {

	getter := new(innerPTGetter)
	getter.init(src)

	getter.Name(subjects.PNameSessionID).GetID(func(value int64) { dst.ID = rbac.SessionID(value) })
	getter.Name(subjects.PNameSessionUUID).GetString(func(value string) { dst.UUID = lang.UUID(value) })

	getter.Name(subjects.PNameUserID).GetID(func(value int64) {
		dst.UserID = rbac.UserID(value)
		dst.Owner = rbac.UserID(value)
	})

	getter.Name(subjects.PNameAvatar).GetString(func(value string) { dst.Avatar = value })
	getter.Name(subjects.PNameNickName).GetString(func(value string) { dst.Nickname = value })
	getter.Name(subjects.PNameUserName).GetString(func(value string) { dst.Username = rbac.UserName(value) })
	getter.Name(subjects.PNameLanguage).GetString(func(value string) { dst.Language = localization.Locale(value) })
	getter.Name(subjects.PNameRoles).GetString(func(value string) { dst.Roles = rbac.RoleNameList(value) })
	getter.Name(subjects.PNameEmail).GetString(func(value string) { dst.Email = rbac.EmailAddress(value) })

	getter.Name(subjects.PNameCreatedAt).GetTime(func(value lang.Time) { dst.CreatedAt = value })
	getter.Name(subjects.PNameUpdatedAt).GetTime(func(value lang.Time) { dst.UpdatedAt = value })
	getter.Name(subjects.PNameDeletedAt).GetTime(func(value lang.Time) { dst.DeletedAt = value })
	getter.Name(subjects.PNameNotBefore).GetTime(func(value lang.Time) { dst.StartedAt = value })
	getter.Name(subjects.PNameNotAfter).GetTime(func(value lang.Time) { dst.ExpiredAt = value })

	getter.Name(subjects.PNameAuthenticated).GetBool(func(value bool) { dst.Authenticated = value })

	getter.GetExt(func(value map[string]string) {
		dst.Properties = value
	})

	return nil
}

func ConvertD2P(src *DTO, dst properties.Table) error {

	setter := new(innerPTSetter)
	setter.init(dst)
	setter.PutMap(src.Properties)

	setter.Name(subjects.PNameSessionID).PutID(int64(src.ID))
	setter.Name(subjects.PNameSessionUUID).PutString(src.UUID.String())

	setter.Name(subjects.PNameUserID).PutID(int64(src.Owner))
	setter.Name(subjects.PNameUserName).PutString(src.Username.String())

	setter.Name(subjects.PNameCreatedAt).PutTime(src.CreatedAt)
	setter.Name(subjects.PNameUpdatedAt).PutTime(src.UpdatedAt)
	setter.Name(subjects.PNameDeletedAt).PutTime(src.DeletedAt)
	setter.Name(subjects.PNameNotBefore).PutTime(src.StartedAt)
	setter.Name(subjects.PNameNotAfter).PutTime(src.ExpiredAt)

	setter.Name(subjects.PNameAvatar).PutString(src.Avatar)
	setter.Name(subjects.PNameEmail).PutString(src.Email.String())
	setter.Name(subjects.PNameNickName).PutString(src.Nickname)
	setter.Name(subjects.PNameLanguage).PutString(src.Language.String())
	setter.Name(subjects.PNameRoles).PutString(string(src.Roles))

	setter.Name(subjects.PNameAuthenticated).PutBool(src.Authenticated)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type innerPTSetter struct {
	table properties.Table
	name  string
}

func (inst *innerPTSetter) init(pt properties.Table) {
	if pt == nil {
		pt = properties.NewTable(nil)
	}
	inst.table = pt
}

func (inst *innerPTSetter) Name(pName subjects.PropertyName) *innerPTSetter {
	inst.name = string(pName)
	return inst
}

func (inst *innerPTSetter) innerPutString(value string) {
	name := inst.name
	inst.table.SetProperty(name, value)
}

func (inst *innerPTSetter) PutString(value string) {
	if value == "" {
		return
	}
	inst.innerPutString(value)
}

func (inst *innerPTSetter) PutID(value int64) {
	if value <= 0 {
		return
	}
	const base = 10
	str := strconv.FormatInt(value, base)
	inst.innerPutString(str)
}

func (inst *innerPTSetter) PutBool(value bool) {

	str := strconv.FormatBool(value)
	inst.innerPutString(str)
}

func (inst *innerPTSetter) PutTime(value lang.Time) {
	n := value.Int()
	if n != 0 {
		const base = 10
		str := strconv.FormatInt(n, base)
		inst.innerPutString(str)
	}
}

func (inst *innerPTSetter) PutTime2t(value time.Time) {
	lt := lang.NewTime(value)
	inst.PutTime(lt)
}

func (inst *innerPTSetter) PutTimeMS(value lang.Milliseconds) {
	num := int64(value)
	if num != 0 {
		const base = 10
		str := strconv.FormatInt(num, base)
		inst.innerPutString(str)
	}
}

func (inst *innerPTSetter) PutMap(src properties.Map) {
	if src == nil {
		return
	}
	tmp := src.Export(nil)
	dst := inst.table
	dst.Import(tmp)
}

////////////////////////////////////////////////////////////////////////////////

type innerPTGetter struct {
	table   properties.Table
	key     string
	keylist []string
}

func (inst *innerPTGetter) init(src properties.Table) {
	if src == nil {
		src = properties.NewTable(nil)
	}
	inst.table = src
}

func (inst *innerPTGetter) Name(name subjects.PropertyName) *innerPTGetter {
	k := string(name)
	inst.key = k
	inst.keylist = append(inst.keylist, k)
	return inst
}

func (inst *innerPTGetter) GetExt(fn func(value map[string]string)) {

	src := inst.table.Export(nil)
	dst := make(map[string]string)
	ignoreTab := make(map[string]bool)
	ignoreList := inst.keylist

	for _, k := range ignoreList {
		ignoreTab[k] = true
	}

	for k, v := range src {
		if ignoreTab[k] {
			continue
		}
		dst[k] = v
	}

	fn(dst)
}

func (inst *innerPTGetter) GetString(fn func(value string)) {
	str, err := inst.innerGetString()
	if err == nil {
		fn(str)
	}
}

func (inst *innerPTGetter) GetID(fn func(value int64)) {
	n, err := inst.innerGetInt64()
	if err == nil {
		fn(n)
	}
}

func (inst *innerPTGetter) GetTime(fn func(value lang.Time)) {
	n, err := inst.innerGetInt64()
	if err == nil {
		fn(lang.Time(n))
	}
}

func (inst *innerPTGetter) GetInt64(fn func(value int64)) {
	n, err := inst.innerGetInt64()
	if err == nil {
		fn(n)
	}
}

func (inst *innerPTGetter) GetBool(fn func(value bool)) {
	str, err := inst.innerGetString()
	if err == nil {
		b, err := strconv.ParseBool(str)
		if err == nil {
			fn(b)
		}
	}
}

func (inst *innerPTGetter) innerGetString() (string, error) {
	name := inst.key
	value := inst.table.GetProperty(name)
	if value == "" {
		return "", fmt.Errorf("no property with name: '%s'", name)
	}
	return value, nil
}

func (inst *innerPTGetter) innerGetInt64() (int64, error) {
	const (
		base  = 10
		bsize = 64
	)
	str, err := inst.innerGetString()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(str, base, bsize)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
