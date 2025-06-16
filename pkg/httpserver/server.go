package httpserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_defaultAddr = ":80"
)

type Server struct {
	App    *http.Server
	Router *mux.Router
	notify chan error
}

func New(opts ...Option) *Server {
	r := mux.NewRouter()

	s := &Server{
		Router: r,
		App: &http.Server{
			Addr:    _defaultAddr,
			Handler: r,
		},
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.App.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.App.Shutdown(ctx)
}
