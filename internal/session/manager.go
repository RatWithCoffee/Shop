package session

import (
	"context"
	"net/http"
)

type Session struct {
	Id uint
}

// interface for removing import cycle
type UserInterface interface {
	GetId() int
}

type SessionManagerInterface interface {
	Check(*http.Request) (*Session, error)
	Create(http.ResponseWriter, UserInterface) error
}

const SESSION_CTX_KEY = "SESSION"

func SessionFromCtx(ctx context.Context) (*Session, bool) {
	sesion, ok := ctx.Value(SESSION_CTX_KEY).(*Session)
	return sesion, ok
}

func AuthMiddleware(next http.Handler, sm SessionManagerInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sm.Check(r)
		if err == nil {
			r = r.WithContext(context.WithValue(r.Context(), SESSION_CTX_KEY, session))
		}
		next.ServeHTTP(w, r)
	})
}
