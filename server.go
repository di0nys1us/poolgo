package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getSessions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sessions", getSessions)

	log.Fatal(http.ListenAndServe(":3001", handlers.CombinedLoggingHandler(os.Stdout, router)))
}
