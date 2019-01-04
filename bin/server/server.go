package server

import (
	"fmt"

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
	out := Server{}
	out.degradation = degradation

	return nil
}

//StartServer create routes and execute http.listenAndServe
func (s *Server) StartServer() error {

	return nil
}
