package repositories

import "providers/models"

type DefaultCustomerRepository struct {
	customers []*models.Customer
	currentID int
}

func NewDefaultCustomerRepository() CustomerRepository {
	r := &DefaultCustomerRepository{
		customers: make([]*models.Customer, 0),
	}

	r.Create("José")

	return r
}

func (r *DefaultCustomerRepository) getNextID() int {
	r.currentID++
	return r.currentID
}

func (r *DefaultCustomerRepository) FindAll() []*models.Customer {
	return r.customers
}

func (r *DefaultCustomerRepository) Find(id int) *models.Customer {
	for _, customer := range r.customers {
		if customer.ID == id {
			return customer
		}
	}

	return nil
}

func (r *DefaultCustomerRepository) Create(name string) *models.Customer {
	customer := &models.Customer{
		ID:   r.getNextID(),
		Name: name,
	}

	r.customers = append(r.customers, customer)

	return customer
}

func (r *DefaultCustomerRepository) Update(id int, name string) *models.Customer {
	customer := r.Find(id)

	if customer == nil {
		return nil
	}

	customer.Name = name

	return customer
}

func (r *DefaultCustomerRepository) Delete(id int) bool {
	for i, customer := range r.customers {
		if customer.ID == id {
			r.customers = append(r.customers[:i], r.customers[i+1:]...)
			return true
		}
	}

	return false
}
