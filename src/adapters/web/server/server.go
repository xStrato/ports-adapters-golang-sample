package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xStrato/ports-adapters-go-sample/src/adapters/web/handlers"
	"github.com/xStrato/ports-adapters-go-sample/src/application/interfaces"
)

type WebServer struct {
	interfaces.IProductService
}

func NewWebServer(s interfaces.IProductService) *WebServer {
	return &WebServer{s}
}

func (w *WebServer) Serve() {
	router := mux.NewRouter()
	middleware := negroni.New(negroni.NewLogger())

	handlers.MakeProductHandler(router, middleware, w.IProductService)
	http.Handle("/", router)

	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":5050",
		Handler:           nil,
		ErrorLog:          log.New(os.Stderr, " log: ", log.Lshortfile),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
