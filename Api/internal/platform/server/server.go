package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuell20/FruitTracker/internal/platform/server/handler/product"
	"github.com/samuell20/FruitTracker/internal/platform/server/handler/user"
	"github.com/samuell20/FruitTracker/internal/platform/server/middleware/logging"
	"github.com/samuell20/FruitTracker/internal/platform/server/middleware/recovery"
	"github.com/samuell20/FruitTracker/kit/command"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	query_services  map[string]interface{}
	shutdownTimeout time.Duration
	commandBus      command.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus, queryServices map[string]interface{}) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		query_services:  queryServices,
		shutdownTimeout: shutdownTimeout,
		commandBus:      commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.Use(recovery.Middleware(), logging.Middleware())

	//Product handlers
	s.engine.GET("/products", product.GetAllHandler(s.query_services["products"]))
	s.engine.POST("/products", product.PostHandler(s.commandBus))
	s.engine.GET("/products/{id}", product.GetHandler(s.query_services["products"]))
	s.engine.PUT("/products/{id}", product.PutHandler(s.commandBus))
	//User handlers
	s.engine.GET("/users", user.GetAllHandler(s.query_services["users"]))
	s.engine.POST("/users", user.PostHandler(s.commandBus))
	s.engine.GET("/users/{id}", user.GetHandler(s.query_services["users"]))
	s.engine.PUT("/users/{id}", user.PutHandler(s.commandBus))
	//Order handlers
	s.engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "ROUTE_NOT_FOUND", "message": "Route not found"})
	})
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()
	log.Println("Servidor ejecutado")
	<-ctx.Done()
	log.Println("Servidor esperando")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)

}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
