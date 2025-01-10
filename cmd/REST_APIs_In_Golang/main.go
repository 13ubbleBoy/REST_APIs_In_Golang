package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/13ubbleBoy/REST_APIs_In_Golang/internal/config"
	"github.com/13ubbleBoy/REST_APIs_In_Golang/internal/http/handlers/student"
	"github.com/13ubbleBoy/REST_APIs_In_Golang/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("database initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))         // for adding a student
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage)) // for getting a student's details by id
	router.HandleFunc("GET /api/students", student.GetList(storage))      // for getting the list of students

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started ", slog.String("at address: ", cfg.Addr))

	// graceful shutdown: jo bhi process chal rahi hai, server band karne par us process ko complete kar k band hoga server.
	done := make(chan os.Signal, 1) // ye Signal type ka channel hai jo "signal: interrupt" espe kaam krega

	// sending Signal to done channel
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // agar 3 defined signal aata hai operating system ki traf se toh done channel ko notify karo.

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()
	<-done // listening to done channel means jab tak done channel k ander koi signal nahi jata tab tak execution finish nahi hoga.

	// logic to stop the server
	slog.Info("shutting down the server")

	// kuch spefici time k baad agar server band nahi hota toh inform karo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // timer agar pass na kare toh server infinitly HANG ho sakte hai
	defer cancel()                                                          // time ko cancel karne k liye

	// err := server.Shutdown(ctx)
	// if err != nil {
	// 	slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	// }
	if err := server.Shutdown(ctx); err != nil { // upper wala or ye same hai.
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
