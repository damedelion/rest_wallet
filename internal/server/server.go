package server

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"github.com/damedelion/rest_wallet/config"
	"github.com/gorilla/mux"
)

type Server struct {
	config *config.Config
	db     *sql.DB
	mux    *mux.Router
}

func NewServer(config *config.Config, db *sql.DB, mux *mux.Router) *Server {
	return &Server{config, db, mux}
}

func (s *Server) Run() {
	s.HandlerRegister()

	server := http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: s.mux,
	}

	/* go func ()  {
		fmt.Println("server is listening on %s", server.Addr)

	} */

	fmt.Printf("server is listening on %s\n", server.Addr)
	server.ListenAndServe()
}
