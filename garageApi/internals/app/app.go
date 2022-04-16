package app

import (
	"context"
	"log"
	"net/http"
	"time"
	"web/garageApi/api"
	"web/garageApi/api/middleware"
	db3 "web/garageApi/internals/app/db"
	"web/garageApi/internals/app/handlers"
	"web/garageApi/internals/app/processors"
	"web/garageApi/internals/cfg"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AppServer struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *AppServer {
	server := new(AppServer)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *AppServer) Serve() {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}

	carsStorage := db3.NewCarStorage(server.db)
	usersStorage := db3.NewUsersStorage(server.db)

	carsProcessor := processors.NewCarsProcessor(carsStorage)
	usersProcessor := processors.NewUsersProcessor(usersStorage)

	userHandler := handlers.NewUsersHandler(usersProcessor)
	carsHandler := handlers.NewCarsHandler(carsProcessor)

	routes := api.CreateRoutes(userHandler, carsHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	return
}

func (server *AppServer) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close()
	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
