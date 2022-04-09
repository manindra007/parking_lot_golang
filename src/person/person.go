package Person

type Person struct {
	Id      int
	Name    string
	Address string
	Mobile  string
	PType   PersonType
}

func createPerson(name string, mobile string, ptype PersonType, address string) *Person {
	return &Person{
		Name:    name,
		Mobile:  mobile,
		PType:   ptype,
		Address: address,
	}
}
