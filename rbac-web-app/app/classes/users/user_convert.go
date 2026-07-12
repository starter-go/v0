package users

import (
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Mobile = src.Mobile
	dst.Language = src.Language
	dst.Avatar = src.Avatar
	dst.DisplayName = src.DisplayName

	dst.Roles = src.Roles.List()

	return nil
}

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Mobile = src.Mobile
	dst.Language = src.Language
	dst.Avatar = src.Avatar
	dst.DisplayName = src.DisplayName

	dst.Roles = src.Roles.Format()

	return nil
}

func ConvertListE2D(src []*Entity, dst []*DTO) ([]*DTO, error) {

	for _, it1 := range src {
		it2 := new(dto.User)
		err := ConvertE2D(it1, it2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, it2)
	}

	return dst, nil
}
