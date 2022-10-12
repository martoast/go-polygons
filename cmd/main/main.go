package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/martoast/go-polygons/pkg/routes"
)

func main() {
	log.Println("API is running!")
	r := mux.NewRouter()
	routes.RegisterPolygonRoutes(r)
	http.Handle("/", r)
	fmt.Println("Listening on http://localhost:8080...")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
