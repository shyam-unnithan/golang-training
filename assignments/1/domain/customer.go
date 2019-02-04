package domain

// Customer structure
type Customer struct {
	ID, Name, Email string
}

//CustomerStore to manage customer
type CustomerStore interface {

	// Create customer information
	Create(Customer) error

	// Delete customer information
	Delete(Customer) error

	// Retrieve customer information by ID
	GetByID(string) Customer

	// Retrieve all customers
	GetAll() map[string]Customer
}
