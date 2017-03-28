package main

import (
	"fmt"
	"log"
	"net/http"
)

type G1Handler struct{}

func (g *G1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Você está no G1!")
}

type GShowHandler struct{}

func (g *GShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Você está no Gshow!")
}

type GeHandler struct{}

func (g *GeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Você está no GE!")
}

func main() {
	addr := "127.0.0.1:8081"
	router := http.NewServeMux()
	router.Handle("/g1", &G1Handler{})
	router.Handle("/gshow", &GShowHandler{})
	router.Handle("/ge", &GeHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
