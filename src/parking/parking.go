package Parking

import (
	"fmt"
	vehicle "parkinglot/src/Vehicle"
)

type Pakring struct {
	parkingid int
	address   string
	// enterance Enterance
	// exit Exit
	Floors []FloorDetail
}

func AddFloor() {

}
func AddSpotToFloor() {
}

func RemveFloor() {

}

func ParkVehicle(veh *vehicle.Vehicle) error {
	for _, Floor := range Floors {
		sp, err := FindFirstEmptySpot(Floor.FloorId, veh.VType)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		ReserveSpot(sp)
		return nil
	}
	return fmt.Errorf("park end")
}

func UnparkVehicle() {

}
