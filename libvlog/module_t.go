package libvlog

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
	"github.com/starter-go/v0/libvlog/gen/main4libvlog"
	"github.com/starter-go/v0/libvlog/gen/test4libvlog"
)

//////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/v0/libvlog"
	theModuleVersion  = "v0.10.0"
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

	mb.Name(theModuleName + "#lib")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	mb.Components(main4libvlog.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	mb.Components(test4libvlog.ExportComponents)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
