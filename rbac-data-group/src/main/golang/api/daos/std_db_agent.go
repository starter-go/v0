package daos

import "github.com/starter-go/libgorm"

type IDatabaseAgent interface {
	libgorm.Agent
}
