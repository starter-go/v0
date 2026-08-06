package api

import "net/http"

const (
	theCoreNS Namespace = "github.com/starter-go/v0/liberr"

	theCoreNoName Name = "no_hyper_error_name"
)

type LibCoreErrorMaker struct {
}

func (inst *LibCoreErrorMaker) innerPrepareInfo(info *HyperErrorInfo) {

	status := http.StatusInternalServerError

	info.Namespace = theCoreNS
	info.StatusCode = status
	info.StatusText = http.StatusText(status)
	info.Name = ""

	// message, status, error, parent ()
	// ns,name
	// uri, hash, code
	ComputeInfoFields(info)

}

func (inst *LibCoreErrorMaker) NewErrorNoNameInSet(name Name) error {
	info := &HyperErrorInfo{
		Name: theCoreNoName,
	}
	inst.innerPrepareInfo(info)
	return NewHyperError(info)
}
