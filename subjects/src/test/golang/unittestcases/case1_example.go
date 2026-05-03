package unittestcases

import "github.com/starter-go/application"

type ExampleCase struct {

	//starter:component

	Enabled bool //starter:inject("${test-case.example.enabled}")

}

// Life implements application.Lifecycle.
func (inst *ExampleCase) Life() *application.Life {

	l := new(application.Life)

	if inst.Enabled {
		l.OnLoop = inst.run
	}

	return l

}

func (inst *ExampleCase) _impl() application.Lifecycle {
	return inst
}

func (inst *ExampleCase) run() error {
	return nil
}
