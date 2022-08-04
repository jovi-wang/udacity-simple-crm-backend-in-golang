package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
)

var CustomerIdNotFound = "Sorry, we can not identify any customer with given ID, please use a valid customerId"
var CustomerIdExist = "Sorry, the customer with given ID exists, please try another ID or update using PUT /customers/{customerId}"

// convert customer dictionary to customer slice
func mapDicToSlice(dictionary map[string]Customer) []Customer {
	customerList := []Customer{}
	for i, v := range dictionary {
		v.ID = i
		customerList = append(customerList, v)
	}
	return customerList
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(mapDicToSlice(customerData))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customerId := mux.Vars(r)["id"]
	if _, ok := customerData[customerId]; ok {
		w.WriteHeader(http.StatusOK)
		customer := customerData[customerId]
		customer.ID = customerId
		json.NewEncoder(w).Encode(customer)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(CustomerIdNotFound)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newEntry := new(Customer)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	customerId, _ := shortid.Generate()
	newEntry.ID = customerId

	if _, ok := customerData[customerId]; ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CustomerIdExist)
	} else {
		customerData[customerId] = *newEntry
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newEntry)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customerId := mux.Vars(r)["id"]

	newEntry := new(Customer)
	newEntry.ID = customerId

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	if _, ok := customerData[customerId]; ok {
		customerData[customerId] = *newEntry
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w)

	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(CustomerIdNotFound)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customerId := mux.Vars(r)["id"]
	if _, ok := customerData[customerId]; ok {
		delete(customerData, customerId)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w)

	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(CustomerIdNotFound)
	}
}
