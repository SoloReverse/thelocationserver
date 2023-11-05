package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	fmt.Println("new server.")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		for k, v := range req.Header {
			fmt.Printf("%v\t%v\n", k, v)
		}
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})

	mux.HandleFunc("/close", func(w http.ResponseWriter, req *http.Request) {

		os.Exit(10)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Listening on server localhost%v\n", s.Addr)

	log.Fatal(s.ListenAndServe())

	defer s.Close()

}
