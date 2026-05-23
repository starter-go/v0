package libjwt

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/libjwt/gen/main4libjwt"
	"github.com/starter-go/v0/libjwt/gen/test4libjwt"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/v0/libjwt"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
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

	mb.Components(main4libjwt.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Components(test4libjwt.ExportComponents)

	mb.Depend(Module())

	mb.Depend(units.Module())
	// mb.Depend(units.ModuleForTest())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
