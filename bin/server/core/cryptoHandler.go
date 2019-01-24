package core

import (
	"fmt"
	"net/http"

	"github.com/michalnov/SovyGo/bin/server/modules/persistance"
	s "github.com/michalnov/SovyGo/bin/server/modules/structures"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//NewKey register new session and create new key-pair for it
func NewKey(w http.ResponseWriter, r *http.Request, p *persistance.Persistance) {
	var env s.Envelop
	err := env.FromEnvelop(r)
	checkErr(err)
	var session s.SessionRequest
	err = session.DecodeSession(env.Body)
	checkErr(err)
	env.Body = p.NewRecord(session.SessionID)
	env.Encryption = false
	w.WriteHeader(http.StatusOK)
	resp, err := env.ToEnvelop()
	checkErr(err)
	fmt.Println(string(resp))
	fmt.Fprintf(w, string(resp))
}
