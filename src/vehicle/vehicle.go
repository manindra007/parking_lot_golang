package Vehicle

import (
	person "parkinglot/src/person"
)

type Vehicle struct {
	RegNo string
	VType Vehicletype
	Color string
	Owner person.Person
}

func createVehicle(regno string, vtype Vehicletype, color string, owner *person.Person) *Vehicle {
	return &Vehicle{
		RegNo: regno,
		VType: vtype,
		Color: color,
		Owner: *owner,
	}
}
