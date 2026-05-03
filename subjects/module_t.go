package subjects

import (
	"embed"

	"github.com/starter-go/application"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/v0/subjects"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
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

func BuildModuleForLib() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	// mb.Components(main4subjects.ExportComponents)
	// mb.Depend(starter.Module())

	return mb
}

func BuildModuleForTest() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	// mb.Components(test4subjects.ExportComponents)
	// mb.Depend(Module())

	return mb
}

////////////////////////////////////////////////////////////////////////////////
// EOF
