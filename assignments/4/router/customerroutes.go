package router

import (
	"assignments/4/controller"

	"github.com/gorilla/mux"
)

//GetRouter ....
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers", controller.PostCustomerHandler).Methods("POST")
	r.HandleFunc("/api/customers", controller.GetCustomerHandler).Methods("GET")
	r.HandleFunc("/api/customers/{id}", controller.GetCustomerByIDHandler).Methods("GET")
	r.HandleFunc("/api/customers/{id}", controller.PutCustomerHandler).Methods("PUT")
	r.HandleFunc("/api/customers/{id}", controller.DeleteCustomerHandler).Methods("DELETE")
	return r
}
