package controller

import (
	"encoding/json"
	"fmt"
	"github.com/adriancarayol/playing-with-golang/pkg/cache"
	"github.com/adriancarayol/playing-with-golang/pkg/services/external/reqres"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	redis *cache.RedisClient
}

func (controller *Controller) Redis() *cache.RedisClient {
	return controller.redis
}

func (controller *Controller) SetRedis(redis *cache.RedisClient) {
	controller.redis = redis
}

func (controller *Controller) GetServiceInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, timestamp, err := controller.redis.GetByKey(id)

	fmt.Println(timestamp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}

	json.NewEncoder(w).Encode(result)
}

func (controller *Controller) UpdateServiceInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	controller.SetResultsFromAPI(id)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	message := fmt.Sprintf("User with id %s retrieved from API (best effort :D)", id)
	w.Write([]byte(message))
}

func (controller *Controller) SetResultsFromAPI(id string) {
	//TODO: Support multiple services
	serviceName := "reqres"

	switch serviceName {
	case "reqres":
		userResult := reqres.User{}
		user := userResult.RetrieveUserGivenId(id)
		b, err := json.Marshal(user)

		if err != nil {
			fmt.Println(err)
			return
		}

		controller.redis.SetValueByKey(id, string(b))
	}
}
