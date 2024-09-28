package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mwives/hexagonal-architecture/adapters/web/handler"
	"github.com/mwives/hexagonal-architecture/app"
)

type WebServer struct {
	Service app.ProductServiceInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (s *WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, n, s.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
