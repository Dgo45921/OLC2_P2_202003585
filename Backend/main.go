package main

import (
	"PY1/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	parserRoutes(router)

	// CORS HANDLER
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // origins
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
	})

	// Envolver el enrutador con el manejador CORS
	handler := corsHandler.Handler(router)

	log.Fatal(http.ListenAndServe(":5000", handler))
}

func parserRoutes(router *mux.Router) {
	router.HandleFunc("/interpreter/parse", controllers.Parse).Methods("POST")
	router.HandleFunc("/interpreter/getCST", controllers.GetCST).Methods("GET")
	router.HandleFunc("/interpreter/getErrors", controllers.GetErrors).Methods("GET")
	router.HandleFunc("/interpreter/getST", controllers.GetSymbolTable).Methods("GET")
}
