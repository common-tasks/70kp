package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	ServiceName string `json:"name"`
	ServiceID   string `json:"id"`
}
type Response struct {
	Services []Service `json:"services"`
}

func main() {
	fmt.Println("starting backend server")
	//create a new router
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")

	fmt.Println("starting server");
	//start and listen to the request
	http.ListenAndServe(":9999",router);

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health check is called")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"API is running")
}
