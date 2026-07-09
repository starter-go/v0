package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security/permissions"
	"github.com/starter-go/v0/subjects"
)

type RbacCheckerController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	PM permissions.Manager //starter:inject("#")
}

// Registration implements libgin.Controller.
func (inst *RbacCheckerController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *RbacCheckerController) route(rp libgin.RouterProxy) error {
	r1 := &libgin.Routing{
		Middleware: true,
		Priority:   9991,
	}
	r1.Handlers = append(r1.Handlers, inst.handle)
	rp.Route(r1)
	return nil
}

func (inst *RbacCheckerController) handle(c *gin.Context) {
	err := inst.checkRoles(c)
	if err != nil {
		code := http.StatusForbidden
		c.AbortWithError(code, err)
	}
}

func (inst *RbacCheckerController) checkRoles(c *gin.Context) error {

	sub, err := subjects.GetCurrent(c)
	if err != nil {
		return err
	}

	gett, err := sub.DoGet()
	if err != nil {
		return err
	}

	subRoles := gett.GetRoles()
	pm := inst.PM
	permWant := new(permissions.Perm)

	permWant.Method = c.Request.Method
	permWant.Path = c.FullPath()

	permHave, err := pm.Find(c, permWant)
	if err != nil {
		return err
	}

	// sub.SetChecker()

	checker := new(innerRbacChecker)
	checker.AcceptRoles(permHave.Roles)
	checker.HaveRoles(subRoles)
	checker.perm = permHave

	return checker.Check()
}

func (inst *RbacCheckerController) _impl() libgin.Controller {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerRbacChecker struct {
	perm  *permissions.Perm
	sgett subjects.Getter
	sub   subjects.Subject
}

func (inst *innerRbacChecker) HasRole(role rbac.RoleName) bool {

	return false
}

func (inst *innerRbacChecker) HaveRoles(roles rbac.RoleNameList) {}

func (inst *innerRbacChecker) AcceptRoles(roles rbac.RoleNameList) {}

func (inst *innerRbacChecker) Check() error {

	p := inst.perm
	s := inst.sub
	sg := inst.sgett

	if p.AcceptAnonym {
		return nil
	}

	if p.AcceptUser && inst.HasRole(rbac.RoleUser) {
		return nil
	}

	if p.AcceptAdmin && inst.HasRole(rbac.RoleAdmin) {
		return nil
	}

	if inst.HasRole(rbac.RoleRoot) {
		return nil
	}

	if p.AcceptOwner && sg.IsAuthenticated() {
		subchecker := new(innerRbacSubjectChecker)
		subchecker.AcceptOwner()
		s.SetChecker(subchecker)
		return nil
	}

	return fmt.Errorf("no permission")
}

////////////////////////////////////////////////////////////////////////////////

type innerRbacSubjectChecker struct {
}

// AcceptAdmin implements subjects.Checker.
func (i *innerRbacSubjectChecker) AcceptAdmin() subjects.Checker {
	panic("unimplemented")
}

// AcceptAnonymous implements subjects.Checker.
func (i *innerRbacSubjectChecker) AcceptAnonymous() subjects.Checker {
	panic("unimplemented")
}

// AcceptAny implements subjects.Checker.
func (i *innerRbacSubjectChecker) AcceptAny() subjects.Checker {
	panic("unimplemented")
}

// AcceptOwner implements subjects.Checker.
func (i *innerRbacSubjectChecker) AcceptOwner() subjects.Checker {
	return i
}

// AcceptRole implements subjects.Checker.
func (i *innerRbacSubjectChecker) AcceptRole(role rbac.RoleName) subjects.Checker {
	panic("unimplemented")
}

// AcceptRoot implements subjects.Checker.
func (i *innerRbacSubjectChecker) AcceptRoot() subjects.Checker {
	panic("unimplemented")
}

// Check implements subjects.Checker.
func (i *innerRbacSubjectChecker) Check() error {
	panic("unimplemented")
}

// CheckObject implements subjects.Checker.
func (i *innerRbacSubjectChecker) CheckObject(ch *subjects.Checking) subjects.Checker {
	panic("unimplemented")
}

////////////////////////////////////////////////////////////////////////////////
