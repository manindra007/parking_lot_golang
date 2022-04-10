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

var display Display

func DisplayBoard(floor FloorDetail) {
	for i, j := range floor.BykeSpotList {
		display.BykeDisplay[i] = j.IsFree
	}
	for i, j := range floor.CarSpotList {
		display.CarDisplay[i] = j.IsFree
	}
}

func Show() {
	for {
		fmt.Println(display.BykeDisplay)
		fmt.Println(display.CarDisplay)
		time.Sleep(1 * time.Minute)
	}
}
