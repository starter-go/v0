package sessions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.NotBefore = src.StartedAt.Time()
	dst.NotAfter = src.ExpiredAt.Time()
	dst.Properties = src.Properties.Text()

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.StartedAt = lang.NewTime(src.NotBefore)
	dst.ExpiredAt = lang.NewTime(src.NotAfter)
	dst.Properties = src.Properties.Map()

	return nil
}

func ConvertListE2D(src []*Entity, dst []*DTO) ([]*DTO, error) {

	for _, o1 := range src {
		o2 := new(DTO)
		err := ConvertE2D(o1, o2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}

	return dst, nil
}
