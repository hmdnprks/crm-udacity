package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers []Customer = []Customer{
	{
		Name:      "John Doe",
		ID:        "001",
		Role:      "customer",
		Email:     "johndoe@gmail.com",
		Phone:     "1234567890",
		Contacted: false,
	},
	{
		Name:      "Jane Doe",
		ID:        "002",
		Role:      "customer",
		Email:     "janedoe@gmail.com",
		Phone:     "9876543210",
		Contacted: false,
	},
	{
		Name:      "John Smith",
		ID:        "003",
		Role:      "customer",
		Email:     "johnsmith@gmail.com",
		Phone:     "1234567890",
		Contacted: false,
	},
}

// The function should return a brief overview of the API (e.g., available endpoints).
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// serve html file
	http.ServeFile(w, r, "index.html")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, customer := range customers {
		if customer.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	// return 404 if customer not found
	w.WriteHeader(http.StatusNotFound)
	// return json message if customer not found
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Customer not found",
	})
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &customer)

	// check if customer already exists
	for _, c := range customers {
		if c.ID == customer.ID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{
				Message: "Customer already exists",
			})
			return
		}
	}

	customers = append(customers, customer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var customer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &customer)

	for i, c := range customers {
		if c.ID == params["id"] {
			customers = append(customers[:i], customers[i+1:]...)
			customers = append(customers, customer)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	// return 404 if customer not found
	w.WriteHeader(http.StatusNotFound)
	// return json message if customer not found
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Customer not found",
	})
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, customer := range customers {
		if customer.ID == params["id"] {
			customers = append(customers[:i], customers[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{
				Message: "Customer deleted",
			})
			return
		}
	}

	// return 404 if customer not found
	w.WriteHeader(http.StatusNotFound)
	// return json message if customer not found
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Customer not found",
	})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	println("Server is listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
