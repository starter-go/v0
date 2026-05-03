package subjects

import (
	"context"
	"fmt"
)

const theAdapterBindingName = "github.com/starter-go/v0/subjects#Adapter/binding"

type Adapter interface {
	GetHolder(c context.Context) (*Holder, error)
}

func GetAdapter(c context.Context) (Adapter, error) {

	const name = theAdapterBindingName
	value := c.Value(name)
	ada, ok := value.(Adapter)

	if !ok {
		return nil, fmt.Errorf("subjects.GetAdapter() : no value with key '%s' in the context", name)
	}

	return ada, nil
}

func SetupAdapter(callback func(bindingName string)) {
	callback(theAdapterBindingName)
}
