package store

import (
	"github.com/Lioric/sessions/user"
)

// ServiceInterface defines the behavior of the session store
type ServiceInterface interface {
	SaveUserSession(userSession *user.Session) error
	DeleteUserSession(sessionID string) error
	FetchValidUserSession(sessionID string) (*user.Session, error)
}
