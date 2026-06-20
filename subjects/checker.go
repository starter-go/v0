package subjects

import (
	"fmt"
	"strings"

	"github.com/starter-go/rbac"
)

type Checking struct {
	Owner rbac.UserID

	Type string

	Target any
}

type Checker interface {
	Check() error

	CheckObject(ch *Checking) Checker

	AcceptRole(role rbac.RoleName) Checker

	AcceptAdmin() Checker
	AcceptRoot() Checker
	AcceptAnonymous() Checker
	AcceptOwner() Checker
}

////////////////////////////////////////////////////////////////////////////////

type innerChecker struct {
	context       *Context
	auth          bool
	user          rbac.UserID
	roles         []rbac.RoleName
	accepts       map[rbac.RoleName]bool
	countOwnerYes int
	countOwnerNo  int
}

func (inst *innerChecker) init(ctx *Context) error {

	getter, err := ctx.Facade.DoGet()
	if err != nil {
		return err
	}

	roles := getter.GetRoles()
	user := getter.GetUserID()
	auth := getter.IsAuthenticated()

	inst.context = ctx
	inst.roles = roles.List()
	inst.user = user
	inst.accepts = make(map[rbac.RoleName]bool)
	inst.auth = auth

	return nil
}

// AcceptAdmin implements Checker.
func (inst *innerChecker) AcceptAdmin() Checker {
	const (
		role = rbac.RoleAdmin
	)
	return inst.AcceptRole(role)
}

// AcceptAnonymous implements Checker.
func (inst *innerChecker) AcceptAnonymous() Checker {
	const (
		role = rbac.RoleAnonym
	)
	return inst.AcceptRole(role)
}

// AcceptOwner implements Checker.
func (inst *innerChecker) AcceptOwner() Checker {
	const (
		role = rbac.RoleOwner
	)
	return inst.AcceptRole(role)
}

// AcceptRole implements Checker.
func (inst *innerChecker) AcceptRole(name rbac.RoleName) Checker {

	str := name.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)

	inst.accepts[rbac.RoleName(str)] = true
	return inst
}

// AcceptRoot implements Checker.
func (inst *innerChecker) AcceptRoot() Checker {
	const (
		role = rbac.RoleRoot
	)
	return inst.AcceptRole(role)
}

// Check implements Checker.
func (inst *innerChecker) Check() error {

	if inst.auth && (inst.user > 0) {
		// is any
		if inst.isAcceptedRole(rbac.RoleAny) {
			return nil
		}
	} else {
		// is anon
		if inst.isAcceptedRole(rbac.RoleAnonym) {
			return nil
		}
		return fmt.Errorf("RBAC:Forbidden")
	}

	const (
		admin = rbac.RoleAdmin
		root  = rbac.RoleRoot
	)

	if inst.isAcceptedRole(admin) {
		if inst.hasRole(admin) {
			return nil
		}
	}

	if inst.isAcceptedRole(root) {
		if inst.hasRole(root) {
			return nil
		}
	}

	if inst.isAcceptedRole(rbac.RoleOwner) {
		y := inst.countOwnerYes
		n := inst.countOwnerNo
		if (y > 0) && (n == 0) {
			return nil
		}
	}

	return fmt.Errorf("RBAC:Forbidden")
}

func (inst *innerChecker) isAcceptedRole(name rbac.RoleName) bool {
	tab := inst.accepts
	if tab == nil {
		return false
	}
	ok := tab[name]
	return ok
}

func (inst *innerChecker) hasRole(name rbac.RoleName) bool {
	all := inst.roles
	for _, have := range all {
		if name == have {
			return true
		}
	}
	return false
}

// CheckObject implements Checker.
func (inst *innerChecker) CheckObject(ch *Checking) Checker {
	panic("unimplemented")
}

func (inst *innerChecker) _impl() Checker {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

func makeNewChecker(ctx *Context) (Checker, error) {
	checker := new(innerChecker)
	err := checker.init(ctx)
	return checker, err
}

////////////////////////////////////////////////////////////////////////////////
// EOF
