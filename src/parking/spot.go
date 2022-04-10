package Parking

import (
	vehicle "parkinglot/src/Vehicle"
)

type Spot struct {
	SpotId   string
	SpotType ParkingType
	IsFree   bool
	Vehicle  vehicle.Vehicle
}

func CreateSpot(spotid string, stype ParkingType) *Spot {
	return &Spot{
		SpotId:   spotid,
		SpotType: stype,
		IsFree:   true,
	}
}

func ReserveSpot(spot *Spot) {
	spot.IsFree = false
}

func FreeSpot(spot *Spot) {
	spot.IsFree = true
}
