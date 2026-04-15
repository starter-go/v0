package users

import (
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

func ConvertE2D(src *entity.User, dst *dto.User) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Mobile = src.Mobile

	dst.Avatar = src.Avatar
	dst.DisplayName = src.DisplayName

	dst.Roles = src.Roles.List()

	return nil
}

func ConvertD2E(src *dto.User, dst *entity.User) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Mobile = src.Mobile

	dst.Avatar = src.Avatar
	dst.DisplayName = src.DisplayName

	dst.Roles = src.Roles.Format()

	return nil
}

func ConvertListE2D(src []*entity.User, dst []*dto.User) ([]*dto.User, error) {

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
