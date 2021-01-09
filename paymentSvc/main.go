package main


import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//Register Handler Function
	handleRequest()
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	// Register API Routes
	//r.HandleFunc("/rest/health", controllers.HealthCheck).Methods(http.MethodGet)
	//r.HandleFunc("/", controllers.Payments).Methods(http.MethodPost)

	//Create HTTP Server
	s := &http.Server{
		Addr: "9090",
		Handler:      r,
		IdleTimeout:  5 * time.Minute, //120
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	//serve up the HTTP Server, make it listen for requests on the specified PORT.
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			log.Fatal("Error starting server on port", s.Addr)
		}
		fmt.Println("Server is up on port", s.Addr)
	}()

	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("Received terminate, graceful shutdown! Signal: ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(tc)

}

