package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Service struct {
	ServiceName string `json:"name"`
	ServiceID   string `json:"id"`
}

func main() {

	//create a new router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/allServices", AllServices)
	router.HandleFunc("/service/{name}", SingleService)
	router.HandleFunc("/services", AddService).Methods("POST")

	var port string = ":9999"
	fmt.Println("starting server")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	//start and listen to the request

	http.ListenAndServe(port, handler)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit : health-check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "health is good")
}

func AllServices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit : AllServices")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetServices())

}
func SingleService(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point hit: specific service")
	vars := mux.Vars(r)
	name := vars["name"]
	for _, service := range GetServices() {
		if service.ServiceName == name {
			json.NewEncoder(w).Encode(service)
		}
	}
}

func AddService(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var service Service
	json.Unmarshal(reqBody, &service)

	services := GetServices()
	services = append(services, service)
	json.NewEncoder(w).Encode(services)
}

func GetServices() []Service {
	services := []Service{}
	var service1 Service
	service1.ServiceID = "1"
	service1.ServiceName = "CurrencyConverter"

	var service2 Service
	service2.ServiceID = "1"
	service2.ServiceName = "TimeConverter"

	services = append(services, service1, service2)
	return services

}
