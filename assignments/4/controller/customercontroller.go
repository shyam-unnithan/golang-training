package controller

import (
	"assignments/4/domain"
	"assignments/4/mapstore"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Cust ...
type Cust struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var controller CustomerController

func init() {
	controller = CustomerController{
		store: mapstore.NewMapStore(),
	}
}

//CustomerController manage customer
type CustomerController struct {
	store domain.CustomerStore
}

//PostCustomerHandler ...
func PostCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var cust Cust
	err := json.NewDecoder(r.Body).Decode(&cust)
	if err != nil {
		log.Fatal(err)
	}
	id := strconv.Itoa(len(controller.store.GetAll()) + 1)

	c := domain.Customer{
		ID:    id,
		Name:  cust.Name,
		Email: cust.Email,
	}
	controller.Add(c)
	j, err := json.Marshal(cust)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetCustomerHandler ...
func GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var custs []Cust
	for _, v := range controller.GetAll() {
		cust := Cust{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		}
		custs = append(custs, cust)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(custs)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetCustomerByIDHandler ...
func GetCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c := controller.GetByID(id)
	cust := Cust{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(cust)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//PutCustomerHandler ...
func PutCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var cust Cust
	err := json.NewDecoder(r.Body).Decode(&cust)

	vars := mux.Vars(r)
	k := vars["id"]

	c := domain.Customer{
		ID:    k,
		Name:  cust.Name,
		Email: cust.Email,
	}
	controller.Update(c)
	j, err := json.Marshal(cust)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(j)
}

//DeleteCustomerHandler ...
func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	c := controller.GetByID(k)
	controller.Delete(c)
	w.WriteHeader(http.StatusNoContent)
}

//Add customer to store
func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("New Customer has been created")
}

//GetByID customer
func (cc CustomerController) GetByID(id string) domain.Customer {
	c, err := cc.store.GetByID(id)
	if err != nil {
		log.Println("Unable to get customer")
	}
	return c
}

//GetAll customers
func (cc CustomerController) GetAll() map[string]domain.Customer {
	return cc.store.GetAll()
}

//Update customer
func (cc CustomerController) Update(c domain.Customer) {
	err := cc.store.Update(c)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("Customer information updated.")
}

//Delete customer
func (cc CustomerController) Delete(c domain.Customer) {
	err := cc.store.Delete(c)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("Customer deleted successfully")
}
