package sessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/v0/rbac-web-app/app/data/entity"
	"github.com/starter-go/v0/rbac-web-app/app/web/dto"
)

func ConvertE2D(src *entity.Session, dst *dto.Session) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.AliveFrom = lang.NewTime(src.AliveFrom)
	dst.AliveTo = lang.NewTime(src.AliveTo)

	return nil
}

func ConvertD2E(src *dto.Session, dst *entity.Session) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.AliveFrom = src.AliveFrom.Time()
	dst.AliveTo = src.AliveTo.Time()

	return nil
}
