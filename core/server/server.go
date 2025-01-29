package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer IServer
}
type IServer interface {
	Start()
}
type HttpServer struct {
	server      *http.Server
	router      *gin.Engine
	environment string
	port        string
	shutDown    chan os.Signal
}

func New(router *gin.Engine, environment string, port string, shutDown chan os.Signal) *Server {
	httpServer := newServer(router, environment, port, shutDown)
	return &Server{
		httpServer: httpServer,
	}
}

func newServer(router *gin.Engine, environment string, port string, down chan os.Signal) *HttpServer {
	return &HttpServer{
		router:      router,
		environment: environment,
		port:        port,
		shutDown:    down,
	}
}

func (server *Server) start() {
	server.httpServer.Start()
}

func (server *Server) Run() {
	server.start()
}

func (server *HttpServer) Start() {
	port := server.port
	if port == "" {
		port = "8080"
	}
	server.server = &http.Server{
		Addr:              port,
		Handler:           server.router,
		ReadHeaderTimeout: 20 * time.Second,
	}
	go func(s *HttpServer) {
		if err := server.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("http server start error:", err)
		}
	}(server)
}
