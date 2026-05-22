package sessions

import (
	"encoding/json"
	"testing"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/localization"
	"github.com/starter-go/v0/subjects"
)

func TestConvertDtoAndProperties(t *testing.T) {

	dt1 := new(DTO)
	dt2 := new(DTO)
	pt1 := properties.NewTable(nil)
	pt2 := properties.NewTable(nil)

	now := lang.Now()
	dt := dt1
	ext := properties.NewTable(nil)

	ext.SetProperty("x", "1.23")
	ext.SetProperty("y", "-45.6")
	ext.SetProperty("z", "0")

	dt.CreatedAt = now
	dt.UpdatedAt = now + 1
	dt.DeletedAt = -2333
	dt.StartedAt = now + 2
	dt.ExpiredAt = now + 3

	dt.ID = 31415926
	dt.UUID = "0000-1111-2222-33333333"
	dt.UserID = 1009527
	dt.Owner = 9527
	dt.Username = "user1"
	dt.Nickname = "Foo .abc Bar"
	dt.Avatar = "http://example.com/a/b/c.txt"
	dt.Language = localization.LocaleZhongCN
	dt.Email = "foo@bar.example.com"
	dt.Roles = "user,admin,root"

	dt.Properties = ext.Export(nil)
	dt.Authenticated = true

	err1 := ConvertD2P(dt1, pt1)
	err2 := ConvertP2D(pt1, dt2)
	err3 := ConvertD2P(dt2, pt2)

	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
	if err3 != nil {
		t.Error(err3)
	}

	const tab = "\t"
	str1, _ := json.MarshalIndent(dt1, tab, tab)
	str2, _ := json.MarshalIndent(dt2, tab, tab)

	t.Logf("pt2 = %v", pt2.Names())
	t.Logf("dt1 = %s", str1)
	t.Logf("dt2 = %s", str2)

}

func TestConvertDtoAndEntity(t *testing.T) {

	dt1 := new(DTO)
	dt2 := new(DTO)
	en1 := new(Entity)
	en2 := new(Entity)

	// pt1 := properties.NewTable(nil)
	// pt2 := properties.NewTable(nil)

	now := lang.Now()
	dt := dt1
	ext := properties.NewTable(nil)

	ext.SetProperty("x", "1.23")
	ext.SetProperty("y", "-45.6")
	ext.SetProperty("z", "0")
	ext.SetProperty(string(subjects.PNameMaxTokenAge), "9876543210")

	dt.CreatedAt = now
	dt.UpdatedAt = now + 1
	dt.DeletedAt = -2333
	dt.StartedAt = now + 2
	dt.ExpiredAt = now + 3

	dt.ID = 31415926
	dt.UUID = "0000-1111-2222-33333333"
	dt.UserID = 1009527
	dt.Owner = 9527
	dt.Username = "user1"
	dt.Nickname = "Foo .abc Bar"
	dt.Avatar = "http://example.com/a/b/c.txt"
	dt.Language = localization.LocaleZhongCN
	dt.Email = "foo@bar.example.com"
	dt.Roles = "user,admin,root"

	dt.Properties = ext.Export(nil)
	dt.Authenticated = true

	err1 := ConvertD2E(dt1, en1)
	err2 := ConvertE2D(en1, dt2)
	err3 := ConvertD2E(dt2, en2)

	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
	if err3 != nil {
		t.Error(err3)
	}

	const tab = "\t"
	str1, _ := json.MarshalIndent(dt1, tab, tab)
	str2, _ := json.MarshalIndent(dt2, tab, tab)
	str4, _ := json.MarshalIndent(en2, tab, tab)

	// t.Logf("pt2 = %v", pt2.Names())

	t.Logf("dt1 = %s", str1)
	t.Logf("dt2 = %s", str2)
	t.Logf("en2 = %s", str4)

}
