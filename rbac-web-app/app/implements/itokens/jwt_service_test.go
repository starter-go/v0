package itokens

import (
	"testing"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

func TestJWTService(t *testing.T) {

	ser := new(JWTokenServiceImpl)
	t1 := new(dto.Token)
	now := lang.Now()

	ser.ConfigKeySecret = "abcdefghijk9876543210"
	ser.ConfigDoVerify = true

	t1.AliveFrom = now - (3600 * 1000) // 1 hour ago
	t1.AliveTo = now + (3600 * 1000)   // 1 hour from now
	t1.SessionID = 2333
	t1.SessionUUID = "aaaa-bbbb-cccc-dddddddd"

	// encode

	jstr1, err := ser.EncodeJWT(t1)
	if err != nil {
		t.Error(err)
		return
	}

	// decode

	t2, err := ser.DecodeJWT(jstr1)
	if err != nil {
		t.Error(err)
		return
	}

	// check result

	ok := true
	if t1.SessionID != t2.SessionID {
		t.Errorf("SessionID mismatch: %v != %v", t1.SessionID, t2.SessionID)
		ok = false
	}
	if t1.SessionUUID != t2.SessionUUID {
		t.Errorf("SessionUUID mismatch: %v != %v", t1.SessionUUID, t2.SessionUUID)
		ok = false
	}
	if t1.AliveFrom != t2.AliveFrom {
		t.Errorf("AliveFrom mismatch: %v != %v", t1.AliveFrom, t2.AliveFrom)
		ok = false
	}
	if t1.AliveTo != t2.AliveTo {
		t.Errorf("AliveTo mismatch: %v != %v", t1.AliveTo, t2.AliveTo)
		ok = false
	}

	if !ok {
		t.Errorf("t1 != t2")
	}
}
