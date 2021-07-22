package main

import (
	"context"

	"github.com/Ghvstcode/RC/controllers"

	"net/http"
	"os"
	"os/signal"
	"time"

	l "github.com/Ghvstcode/RC/utils/logger"

	"github.com/gorilla/mux"
)

func main() {
	//utils.DbCron()
	handleRequest()
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/set", controllers.Set)

	r.Path("/get").Queries("key", "").HandlerFunc(controllers.Get).Methods(http.MethodPost)

	s := &http.Server{
		Addr:         ":4000",
		Handler:      r,
		IdleTimeout:  5 * time.Minute, //120
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	go func() {
		l.InfoLogger.Println("Server is up on port", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			l.ErrorLogger.Println(err)
			l.ErrorLogger.Fatal("Error starting server on port", s.Addr)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.InfoLogger.Println("Received terminate, graceful shutdown! Signal: ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(tc)

}
