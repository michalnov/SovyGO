package core

import (
	"net/http"

	"github.com/michalnov/SovyGo/bin/server/modules/persistance"
	s "github.com/michalnov/SovyGo/bin/server/modules/structures"
)

//NewKey register new session and create new key-pair for it
func NewKey(w *http.ResponseWriter, r *http.Request, p *persistance.Persistance) {
	var session s.SessionRequest
	session, err := s.DecodeSession(r)
	if err != nil {
		panic(err)
	}
	p.NewRecord(session.SessionID)
}
