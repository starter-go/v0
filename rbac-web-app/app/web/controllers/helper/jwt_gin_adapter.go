package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/v0/libjwt/api/jwt"
)

type GinLibjwtAdapter struct {

	//starter:component

	_as func(jwt.AdapterProvider) //starter:as(".")

	MyName      string //starter:inject("${jwt-adapter.libgin.name}")
	MyPriority  int    //starter:inject("${jwt-adapter.libgin.priority}")
	MyEnabled   bool   //starter:inject("${jwt-adapter.libgin.enabled}")
	MyUseCookie bool   //starter:inject("${jwt-adapter.libgin.use-http-cookie}")
	MyUseHeader bool   //starter:inject("${jwt-adapter.libgin.use-http-header}")
	MyUseQuery  bool   //starter:inject("${jwt-adapter.libgin.use-http-query}")

}

// Registration implements jwt.AdapterProvider.
func (inst *GinLibjwtAdapter) Registration() *jwt.AdapterRegistration {

	reg := &jwt.AdapterRegistration{
		Name:     inst.MyName,
		Enabled:  inst.MyEnabled,
		Priority: inst.MyPriority,

		Loader:   inst,
		Provider: inst,
	}
	return reg
}

// Load implements jwt.AdapterLoader.
func (inst *GinLibjwtAdapter) Load(target *jwt.AdapterRegistration) error {
	target.Loader = inst
	target.Adapter = inst
	target.Provider = inst
	return nil
}

// GetToken implements jwt.Adapter.
func (inst *GinLibjwtAdapter) GetToken(acc *jwt.Access) error {

	gc, err := inst.innerGetGinContext(acc)
	if err != nil {
		return err
	}

	// get from (cookie|header|query)

	acc.Text = ""
	inst.innerGetFromHeader(gc, acc)
	if acc.Text == "" {
		inst.innerGetFromQuery(gc, acc)
	}
	if acc.Text == "" {
		inst.innerGetFromCookie(gc, acc)
	}

	// decode

	err = acc.Decode(false)
	if err != nil {
		return err
	}

	return nil
}

// SetToken implements jwt.Adapter.
func (inst *GinLibjwtAdapter) SetToken(acc *jwt.Access) error {

	gc, err := inst.innerGetGinContext(acc)
	if err != nil {
		return err
	}

	// encode
	err = acc.Encode(false)
	if err != nil {
		return err
	}

	// set to (cookie|header)
	inst.innerSetToCookie(gc, acc)
	inst.innerSetToHeader(gc, acc)

	return nil
}

func (inst *GinLibjwtAdapter) innerSetToCookie(gc *gin.Context, acc *jwt.Access) error {

	if !inst.MyUseCookie {
		return nil
	}

	const (
		name = "jwt"
	)

	value := acc.Text.String()
	maxAge := 3600
	path := ""
	domain := ""
	secure := false
	httpOnly := false

	gc.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
	return nil
}

func (inst *GinLibjwtAdapter) innerSetToHeader(gc *gin.Context, acc *jwt.Access) error {

	if !inst.MyUseHeader {
		return nil
	}
	const (
		name = "x-set-jwt"
	)
	value := acc.Text.String()
	gc.Header(name, value)
	return nil
}

func (inst *GinLibjwtAdapter) innerGetFromCookie(gc *gin.Context, acc *jwt.Access) error {

	if !inst.MyUseCookie {
		return nil
	}

	const (
		name = "jwt"
	)

	value, err := gc.Cookie(name)
	if err != nil {
		return err
	}

	acc.Text = jwt.Text(value)
	return nil
}

func (inst *GinLibjwtAdapter) innerGetFromHeader(gc *gin.Context, acc *jwt.Access) error {

	if !inst.MyUseHeader {
		return nil
	}
	const (
		name = "x-jwt"
	)
	value := gc.GetHeader(name)
	if value == "" {
		return fmt.Errorf("no value in HTTP Request Header, name = %s", name)
	}
	acc.Text = jwt.Text(value)
	return nil
}

func (inst *GinLibjwtAdapter) innerGetFromQuery(gc *gin.Context, acc *jwt.Access) error {

	if !inst.MyUseQuery {
		return nil
	}

	const (
		name1 = "jwt"
		name2 = "token"
	)

	value := gc.Query(name1)
	if value != "" {
		acc.Text = jwt.Text(value)
		return nil
	}

	value = gc.Query(name2)
	if value != "" {
		acc.Text = jwt.Text(value)
		return nil
	}

	return nil
}

func (inst *GinLibjwtAdapter) innerGetGinContext(acc *jwt.Access) (*gin.Context, error) {
	cc := acc.Context
	gc, ok := cc.(*gin.Context)
	if !ok {
		return nil, fmt.Errorf("GinLibjwtAdapter: the context.Context is NOT a gin.Context")
	}
	return gc, nil
}

func (inst *GinLibjwtAdapter) _impl() (jwt.AdapterProvider, jwt.AdapterLoader, jwt.Adapter) {
	return inst, inst, inst
}
