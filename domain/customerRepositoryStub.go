package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func NewCustomerRepository() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Felipe", City: "Santiago", Zipcode: "1213123", DateofBirth: "2000-02-02", Status: "1"},
		{Id: "1002", Name: "Andres", City: "Santiago", Zipcode: "1213123", DateofBirth: "2000-02-01", Status: "1"},
	}
	return CustomerRepositoryStub{customer: customers}
}
