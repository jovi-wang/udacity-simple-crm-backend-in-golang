package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Server is starting on port 3000...")

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":3000", router)
}
