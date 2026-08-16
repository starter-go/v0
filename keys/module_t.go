package keys

import (
	"embed"

	"github.com/starter-go/application"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/v0/keys"
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

	mb.Name(theModuleName + "#lib").Version(theModuleVersion).Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	return mb
}

func BuildModuleForTest() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test").Version(theModuleVersion).Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}

////////////////////////////////////////////////////////////////////////////////
// EOF
