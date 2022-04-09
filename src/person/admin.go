package Person

func CreateAdmin(name string, mobile string, address string) {
	createPerson(name, mobile, PersonType(Admin), address)
}
