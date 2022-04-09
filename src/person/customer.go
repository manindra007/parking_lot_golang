package Person

func CreateCustomer(name string, mobile string, address string) *Person {
	return createPerson(name, mobile, PersonType(Customer), address)
}
