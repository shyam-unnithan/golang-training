package mapstore

import (
	"assignments/1/domain"
	"errors"
)

// MapStore for storing customer
type MapStore struct {
	store map[string]domain.Customer
}

// NewMapStore - Factory to create a new MapStore for Customer
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

// Create customer
func (mapStore MapStore) Create(c domain.Customer) error {
	if _, ok := mapStore.store[c.ID]; ok {
		return errors.New("Customer exits")
	}
	mapStore.store[c.ID] = c
	return nil
}

//Update Customer
func (mapStore MapStore) Update(c domain.Customer) error {
	if _, ok := mapStore.store[c.ID]; !ok {
		return errors.New("Failed -> Customer not found")
	}
	mapStore.store[c.ID] = c
	return nil

}

// Delete customer
func (mapStore MapStore) Delete(c domain.Customer) error {
	if _, ok := mapStore.store[c.ID]; !ok {
		return errors.New("Customer not found")
	}
	delete(mapStore.store, c.ID)
	return nil
}

//GetAll customers
func (mapStore MapStore) GetAll() map[string]domain.Customer {
	return mapStore.store
}

//GetByID customers by id
func (mapStore MapStore) GetByID(id string) (c domain.Customer, err error) {
	if _, ok := mapStore.store[id]; !ok {
		var dc domain.Customer
		return dc, errors.New("Customer not found")
	}
	return mapStore.store[id], nil
}
