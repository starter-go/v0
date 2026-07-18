package permissions

import (
	"strings"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
)

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbac.CopyBaseFieldsD2E(src, dst)

	dst.Enabled = src.Enabled
	dst.Method = src.Method
	dst.Path = src.Path
	dst.Priority = src.Priority
	dst.Roles = src.Roles

	uri := innerComputeResUri(src.Method, src.Path)
	dst.URI = lang.URI(uri)

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbac.CopyBaseFieldsE2D(src, dst)

	dst.Enabled = src.Enabled
	dst.Method = src.Method
	dst.Path = src.Path
	dst.Priority = src.Priority
	dst.Roles = src.Roles
	dst.URI = src.URI

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
