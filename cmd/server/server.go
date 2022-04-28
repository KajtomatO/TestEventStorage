package server

import (
	"log"
	"net/http"

	"github.com/KajtomatO/TestEventStorage/internal/app/server_lib"
	"github.com/KajtomatO/TestEventStorage/internal/app/storage_memory"
)

func main() {
	server := &server_lib.DataServer{storage_memory.NewInMemoryDataStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
