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

var Floors []FloorDetail

func CreateFloor(parkingId string) {
	Floors = append(Floors, FloorDetail{
		FloorId:     parkingId + string(len(Floors)+1),
		FloorNumber: len(Floors) + 1,
	})
}

func AddSpot(floorid string, spottype Vehicle.Vehicletype) {
	for i, floor := range Floors {
		if floor.FloorId != floorid {
			continue
		}
		if spottype == Vehicle.Car {
			spot := CreateSpot(floorid+"S"+string(len(floor.CarSpotList)+1), ParkingType(Car))
			Floors[i].CarSpotList = append(Floors[i].CarSpotList, *spot)
			break
		} else if spottype == Vehicle.Byke {
			spot := CreateSpot(floorid+"S"+string(len(floor.BykeSpotList)+1), ParkingType(Byke))
			Floors[i].BykeSpotList = append(Floors[i].BykeSpotList, *spot)
			break
		}
	}
}

func RemoveSpotById(spotid string, spottype Vehicle.Vehicletype) {
	for i, floor := range Floors {
		if floor.FloorId != spotid {
			continue
		}
		if spottype == Vehicle.Byke {
			if err := removespot(spotid, Floors[i].BykeSpotList); err != nil {
				continue
			} else {
				break
			}
		} else if spottype == Vehicle.Car {
			if err := removespot(spotid, Floors[i].CarSpotList); err != nil {
				continue
			} else {
				break
			}
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
func FindFirstEmptySpot(floorid string, spottype Vehicle.Vehicletype) (*Spot, error) {
	var spot *Spot
	var err error
	var floor FloorDetail
	for _, f := range Floors {
		if f.FloorId == floorid {
			floor = f
		}
	}
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
