package Parking

import (
	"fmt"
	Vehicle "parkinglot/src/Vehicle"
)

type FloorDetail struct {
	FloorId      int
	FloorNumber  int
	CarSpotList  []Spot
	BykeSpotList []Spot
}

var Floor FloorDetail

func AddSpot(spottype Vehicle.Vehicletype) {
	spot := CreateSpot()
	if spottype == Vehicle.Car {
		Floor.CarSpotList = append(Floor.CarSpotList, *spot)
	} else if spottype == Vehicle.Byke {
		Floor.BykeSpotList = append(Floor.BykeSpotList, *spot)
	}
}

func RemoveSpotById(spotid int, spottype Vehicle.Vehicletype) {
	if spottype == Vehicle.Byke {
		removespot(spotid, Floor.BykeSpotList)
	}
}

func removespot(spotid int, spots []Spot) error {
	for i, j := range spots {
		if spotid == j.SpotId {
			spots = append(spots[0:i-1], spots[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Spot not found")
}
func FindFirstEmptySpot(spottype Vehicle.Vehicletype) (*Spot, error) {
	var spot *Spot
	var err error
	if spottype == Vehicle.Car {
		spot, err = findspot(Floor.CarSpotList)
		if err != nil {
			return spot, err
		}
	} else {
		spot, err = findspot(Floor.BykeSpotList)
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
	return nil, fmt.Errorf("no spot found")
}
