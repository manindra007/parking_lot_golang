package Person

func CreateEmloyee(name string, mobile string, address string) {
	createPerson(name, mobile, PersonType(Employee), address)
}
