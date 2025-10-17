package sessions

import "context"

type Service interface {
	FindSession(c context.Context, id SessionID) (*SessionDTO, error)

	InsertSession(c context.Context, item *SessionDTO) (*SessionDTO, error)

	UpdateSession(c context.Context, id SessionID, item *SessionDTO) (*SessionDTO, error)

	RemoveSession(c context.Context, id SessionID) error
}
