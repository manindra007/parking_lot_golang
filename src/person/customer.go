package Person

func CreateCustomer(name string, mobile string, address string) {
	createPerson(name, mobile, PersonType(Customer), address)
}
