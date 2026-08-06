package pack1

import "testing"

func TestDemoErrorSet(t *testing.T) {

	var es Demo1ErrorSet

	err := es.MakeError2("foo", 1000, false, 3.14)

	t.Log(err.Error())

}
