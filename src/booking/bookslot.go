package Booking

import (
	"fmt"
	vehicle "parkinglot/src/Vehicle"
	park "parkinglot/src/parking"
	person "parkinglot/src/person"
)

func BookSlot() {
	park.CreateParkingLot()
	p := park.GetParking()
	p.AddFloorInParking()
	p.AddFloorInParking()
	park.AddSpotToFloor(0)
	u1 := person.CreateCustomer("manindra", "8299661294", "")
	v1 := vehicle.CreateCar("UP53AP0001", "grey", u1)
	err := park.ParkVehicle(v1)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println("parked")
	}
}
