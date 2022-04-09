package Vehicle

import Person "parkinglot/src/person"

func CreateCar(regno string, color string, owner Person.Person) *Vehicle {
	return createVehicle(regno, Vehicletype(Car), color, owner)
}
