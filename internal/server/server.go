package server

import (
	"net/http"
	"os"

	"github.com/AlexDillz/Calc_server_yandex/internal/calculator"
	"github.com/gorilla/mux"
)

type Server struct {
	Config *Config
}

type Config struct {
	Addr string
}

func New() *Server {
	return &Server{
		Config: &Config{
			Addr: getPort(),
		},
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func (s *Server) Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/calculate", calculator.Handler).Methods("POST")

	http.Handle("/", r)
	return http.ListenAndServe(s.Config.Addr, nil)
}
