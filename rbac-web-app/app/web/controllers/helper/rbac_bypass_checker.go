package helper

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/subjects"
)

type innerBypassChecker struct {
}

// AcceptAdmin implements [subjects.Checker].
func (inst *innerBypassChecker) AcceptAdmin() subjects.Checker {
	return inst
}

// AcceptAnonymous implements [subjects.Checker].
func (inst *innerBypassChecker) AcceptAnonymous() subjects.Checker {
	return inst
}

// AcceptAny implements [subjects.Checker].
func (inst *innerBypassChecker) AcceptAny() subjects.Checker {
	return inst
}

// AcceptOwner implements [subjects.Checker].
func (inst *innerBypassChecker) AcceptOwner() subjects.Checker {
	return inst
}

// AcceptRole implements [subjects.Checker].
func (inst *innerBypassChecker) AcceptRole(role rbac.RoleName) subjects.Checker {
	return inst
}

// AcceptRoot implements [subjects.Checker].
func (inst *innerBypassChecker) AcceptRoot() subjects.Checker {
	return inst
}

// Check implements [subjects.Checker].
func (inst *innerBypassChecker) Check() error {
	return nil
}

// CheckDTO implements [subjects.Checker].
func (inst *innerBypassChecker) CheckDTO(ref rbac.DTORef) subjects.Checker {
	return inst
}

// CheckEntity implements [subjects.Checker].
func (inst *innerBypassChecker) CheckEntity(ref rbac.EntityRef) subjects.Checker {
	return inst
}

// CheckObject implements [subjects.Checker].
func (inst *innerBypassChecker) CheckObject(ch *subjects.Checking) subjects.Checker {
	return inst
}

func (inst *innerBypassChecker) _impl() subjects.Checker {
	return inst
}
