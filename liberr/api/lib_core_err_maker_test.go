package api

import "testing"

func TestCoreErrorMaker(t *testing.T) {

	mkr := new(LibCoreErrorMaker)
	err := mkr.NewErrorNoNameInSet("foo")

	t.Log(err.Error())

}
