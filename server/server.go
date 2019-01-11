package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Server structure that hold all parts of application
type Server struct {
	r           *mux.Router
	degradation chan int
}

//SetupServer prepare new server structure
func (s *Server) SetupServer(degradation chan int) error {
	fmt.Println("Creating server")
	s.degradation = degradation
	return nil
}

//StartServer create routes and execute http.listenAndServe
func (s *Server) StartServer() error {
	s.r = mux.NewRouter()
	s.r.HandleFunc("/", homeHandler)
	s.r.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("shutdown"))
		s.degradation <- 0
	})
	s.r.HandleFunc("/hello", notImplemented)
	http.Handle("/", s.r)
	http.ListenAndServe(":1122", s.r)
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello at Server Home"))
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("not implemented yet"))
}
