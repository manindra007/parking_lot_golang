package Booking

import (
	"fmt"
	vehicle "parkinglot/src/Vehicle"
	park "parkinglot/src/parking"
	person "parkinglot/src/person"
)

func BookSlot() {
	u1 := person.CreateCustomer("manindra", "8299661294", "")
	fmt.Println("here  1", *u1)
	v1 := vehicle.CreateCar("UP53AP0001", "grey", u1)
	fmt.Println("here 2", *v1)
	err := park.ParkVehicle(v1)
	if err != nil {
		fmt.Println(err.Error())
	}
}
