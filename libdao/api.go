package libdao

import "github.com/starter-go/v0/libdao/api/libdaoapi"

type DAO = libdaoapi.DAO

type DaoRegistration = libdaoapi.DaoRegistration

type DaoHolder[T DAO] = libdaoapi.DaoHolder[T]
