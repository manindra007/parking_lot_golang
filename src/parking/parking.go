package Parking

import (
	"fmt"
	vehicle "parkinglot/src/Vehicle"
)

type Parking struct {
	Parkingid string
	Address   string
	// enterance Enterance
	// exit Exit
	Floors []FloorDetail
}

var ParkingLot Parking

func CreateParkingLot() {
	ParkingLot.Parkingid = "abc"
}
func GetParking() *Parking {
	return &ParkingLot
}

func (p Parking) AddFloorInParking() {
	ParkingLot.Floors = append(ParkingLot.Floors, *CreateFloor(ParkingLot.Parkingid + "F"))
}
func AddSpotToFloor(index int) {
	AddSpot(&ParkingLot.Floors[index], vehicle.Car)
}

func RemveFloor() {

}

func ParkVehicle(veh *vehicle.Vehicle) error {
	for _, Floor := range ParkingLot.Floors {
		sp, err := FindFirstEmptySpot(Floor, veh.VType)
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
