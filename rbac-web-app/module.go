package rbacwebapp

import (
	"embed"

	"github.com/starter-go/application"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/v0/rbac-web-app"
	theModuleVersion  = "v0.0.2"
	theModuleRevision = 2
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath    = "src/main/resources"
	theTestModuleResPath    = "src/test/resources"
	theTestWebCommonResPath = "src/web-test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/web-test/resources"
var theTestWebCommonModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func NewMainModuleBuilder() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	return mb
}

func NewTestModuleBuilder() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test-client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	return mb
}

func NewTestWebServerModuleBuilder() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test-web-server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestWebCommonModuleResFS, theTestWebCommonResPath)

	return mb
}

func NewTestWebClientModuleBuilder() *application.ModuleBuilder {

	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test-web-client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestWebCommonModuleResFS, theTestWebCommonResPath)

	return mb
}

////////////////////////////////////////////////////////////////////////////////
// EOF
