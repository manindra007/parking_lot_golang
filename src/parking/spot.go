package Parking

import (
	vehicle "parkinglot/src/Vehicle"
)

type Spot struct {
	SpotId   int
	SpotType ParkingType
	IsFree   bool
	Vehicle  vehicle.Vehicle
}

func CreateSpot() *Spot {
	return &Spot{}
}

func ReserveSpot(spot *Spot) {
	spot.IsFree = false
}

func FreeSpot(spot *Spot) {
	spot.IsFree = true
}
