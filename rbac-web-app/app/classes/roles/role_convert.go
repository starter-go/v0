package roles

import (
	"github.com/starter-go/rbac"
)

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbac.CopyBaseFieldsD2E(src, dst)

	dst.Name = src.Name
	dst.Description = src.Description

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbac.CopyBaseFieldsE2D(src, dst)

	dst.Name = src.Name
	dst.Description = src.Description

	return nil
}

func ConvertListE2D(src []*Entity, dst []*DTO) ([]*DTO, error) {

	for _, it1 := range src {
		it2 := new(DTO)
		err := ConvertE2D(it1, it2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, it2)
	}

	return dst, nil
}
