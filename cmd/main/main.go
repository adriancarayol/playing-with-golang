package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/adriancarayol/playing-with-golang/pkg/cache"
	"github.com/adriancarayol/playing-with-golang/pkg/services/controller"
)

func main() {
	redisDB, err := cache.InitRedisClient()

	if err != nil {
		log.Fatal(err)
	}

	apiController := &controller.Controller{}
	apiController.SetRedis(redisDB)

	router := mux.NewRouter()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	router.HandleFunc("/service/{id}", apiController.GetServiceInfo).Methods("GET")
	router.HandleFunc("/update/{id}", apiController.UpdateServiceInfo).Methods("GET")
	http.ListenAndServe(":1800", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))

}
