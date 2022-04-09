package Vehicle

import Person "parkinglot/src/person"

func CreateByke(regno string, color string, owner Person.Person) *Vehicle {
	return createVehicle(regno, Vehicletype(Byke), color, owner)
}
