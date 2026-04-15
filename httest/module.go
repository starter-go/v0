package httest

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/v0/httest/gen/main4httest"
	"github.com/starter-go/v0/httest/gen/test4httest"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/v0/httest"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 2
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func Module() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Components(main4httest.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Components(test4httest.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
