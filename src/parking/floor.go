package Parking

import (
	"fmt"
	Vehicle "parkinglot/src/Vehicle"
)

type FloorDetail struct {
	FloorId      string
	FloorNumber  int
	CarSpotList  []Spot
	BykeSpotList []Spot
}

func CreateFloor(floorid string) *FloorDetail {
	return &FloorDetail{
		FloorId: floorid,
	}
}

func AddSpot(Floor *FloorDetail, spottype Vehicle.Vehicletype) {
	if spottype == Vehicle.Car {
		spot := CreateSpot(Floor.FloorId+"S", ParkingType(Car))
		Floor.CarSpotList = append(Floor.CarSpotList, *spot)
	} else if spottype == Vehicle.Byke {
		spot := CreateSpot(Floor.FloorId+"S"+string(len(Floor.BykeSpotList)+1), ParkingType(Byke))
		Floor.BykeSpotList = append(Floor.BykeSpotList, *spot)
	}
}

func RemoveSpotById(floor FloorDetail, spotid string, spottype Vehicle.Vehicletype) {
	if spottype == Vehicle.Byke {
		if err := removespot(spotid, floor.BykeSpotList); err != nil {
		}
	} else if spottype == Vehicle.Car {
		if err := removespot(spotid, floor.CarSpotList); err != nil {
		}
	}
}

func removespot(spotid string, spots []Spot) error {
	for i, j := range spots {
		if spotid == j.SpotId {
			spots = append(spots[0:i-1], spots[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Spot not found")
}
func FindFirstEmptySpot(floor FloorDetail, spottype Vehicle.Vehicletype) (*Spot, error) {
	var spot *Spot
	var err error
	if spottype == Vehicle.Car {
		spot, err = findspot(floor.CarSpotList)
		if err != nil {
			return spot, err
		}
	} else {
		spot, err = findspot(floor.BykeSpotList)
		if err != nil {
			return spot, err
		}
	}
	return spot, nil
}

func findspot(spots []Spot) (*Spot, error) {
	for _, spot := range spots {
		if spot.IsFree == true {
			return &spot, nil
		}
	}
	return nil, fmt.Errorf("Spot not found")
}
