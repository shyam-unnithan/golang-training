package main

import (
	"assignments/1/domain"
	"assignments/1/mapstore"
	"fmt"
)

//CustomerController manage customer
type CustomerController struct {
	store domain.CustomerStore
}

//Add customer to store
func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
}

//GetByID customer
func (cc CustomerController) GetByID(id string) domain.Customer {
	return cc.store.GetByID(id)
}

//GetAll customers
func (cc CustomerController) GetAll() {
	fmt.Println(cc.store.GetAll())
}

//Delete customer
func (cc CustomerController) Delete(c domain.Customer) {
	err := cc.store.Delete(c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer deleted successfully")
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}

	customer := domain.Customer{
		ID:   "cust101",
		Name: "Test",
	}
	controller.Add(customer)
	controller.GetAll()
	customer = controller.GetByID("cust101")
	controller.Delete(customer)
	controller.Add(customer)
}
