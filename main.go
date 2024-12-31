package main

import (
	"fmt"
	"time"

	"net/http"
)

// Simple counter server.
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
	fmt.Println(ctr.n)
}

func main() {

	httpCustomHandler := new(Counter)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        httpCustomHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println(s.ListenAndServe())
	// go mod tidy && go run . - basic commands
}
