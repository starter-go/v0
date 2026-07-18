package tables

import "github.com/starter-go/rbac"

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbac.CopyBaseFieldsD2E(src, dst)

	dst.Description = src.Description
	dst.GroupName = src.GroupName
	dst.GroupURI = src.GroupURI
	dst.Label = src.Label
	dst.Name = src.Name
	dst.TableURI = src.TableURI

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbac.CopyBaseFieldsE2D(src, dst)

	dst.Description = src.Description
	dst.GroupName = src.GroupName
	dst.GroupURI = src.GroupURI
	dst.Label = src.Label
	dst.Name = src.Name
	dst.TableURI = src.TableURI

	return nil
}
