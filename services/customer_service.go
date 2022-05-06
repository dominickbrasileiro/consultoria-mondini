package services

import (
	"providers/models"
	"providers/repositories"
)

type CustomerService struct {
	repository repositories.CustomerRepository
}

func NewCustomerService(r repositories.CustomerRepository) *CustomerService {
	return &CustomerService{r}
}

func (r *CustomerService) FindAll() []*models.Customer {
	return r.repository.FindAll()
}

func (r *CustomerService) Find(id int) *models.Customer {
	return r.repository.Find(id)
}

func (r *CustomerService) Create(name string) *models.Customer {
	return r.repository.Create(name)
}

func (r *CustomerService) Update(id int, name string) *models.Customer {
	return r.repository.Update(id, name)
}

func (r *CustomerService) Delete(id int) bool {
	return r.repository.Delete(id)
}
