package permissions

import (
	"strings"

	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Method = src.Method
	dst.Path = src.Path
	dst.AcceptRoles = src.Roles
	dst.Enabled = src.Enabled

	dst.Resource = innerComputeResUri(src.Method, src.Path)

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Method = src.Method
	dst.Path = src.Path
	dst.Roles = src.AcceptRoles
	dst.Enabled = src.Enabled

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

////////////////////////////////////////////////////////////////////////////////

func innerComputeResUri(method, path string) string {

	b := new(strings.Builder)

	b.WriteString("uri://perm:0")
	b.WriteString(strings.TrimSpace(path))
	b.WriteString("#")
	b.WriteString(strings.TrimSpace(method))

	str := b.String()
	return strings.ToLower(str)
}
