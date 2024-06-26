package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	log.Printf("Server is running on %s", *addr)
	http.ListenAndServe(*addr, routes())
}