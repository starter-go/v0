package main4httest
import (
    p0ef6f2938 "github.com/starter-go/application"
    p25d518c24 "github.com/starter-go/v0/httest/libs/httestrunner"
     "github.com/starter-go/application"
)

// type p25d518c24.RunnerCom in package:github.com/starter-go/v0/httest/libs/httestrunner
//
// id:com-25d518c242143dba-httestrunner-RunnerCom
// class:
// alias:
// scope:singleton
//
type p25d518c242_httestrunner_RunnerCom struct {
}

func (inst* p25d518c242_httestrunner_RunnerCom) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-25d518c242143dba-httestrunner-RunnerCom"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p25d518c242_httestrunner_RunnerCom) new() any {
    return &p25d518c24.RunnerCom{}
}

func (inst* p25d518c242_httestrunner_RunnerCom) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p25d518c24.RunnerCom)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)


    return nil
}


func (inst*p25d518c242_httestrunner_RunnerCom) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


