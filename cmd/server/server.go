package server

import (
	"log"
	"net/http"
)

func main() {
	server := &DataServer{NewInMemoryDataStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
