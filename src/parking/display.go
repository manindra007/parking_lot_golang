package Parking

import (
	"fmt"
	"time"
)

type DisplayCar bool
type DisplayByke bool

type Display struct {
	CarDisplay  []bool
	BykeDisplay []bool
}

var display []Display

func DisplayBoard(floor FloorDetail) {
	for i, j := range floor.BykeSpotList {
		display[floor.FloorId].BykeDisplay[i] = j.IsFree
	}
	for i, j := range floor.CarSpotList {
		display[floor.FloorId].CarDisplay[i] = j.IsFree
	}
}

func Show() {
	for {
		fmt.Println(display[Floor.FloorId].BykeDisplay)
		fmt.Println(display[Floor.FloorId].CarDisplay)
		time.Sleep(1 * time.Minute)
	}
}
