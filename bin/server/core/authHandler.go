package core

import (
	"net/http"

	"github.com/michalnov/SovyGo/bin/server/modules/persistance"
	s "github.com/michalnov/SovyGo/bin/server/modules/structures"
)

//LoginHandler handle proces of user login
func LoginHandler(w *http.ResponseWriter, r *http.Request, p *persistance.Persistance) {
	var req s.LoginRequest
	req, err := s.DecodeLogin(r)
	if err != nil {

		panic(err)
	}

}
