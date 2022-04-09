package Parking

type Spot struct {
	SpotId   int
	SpotType ParkingType
	IsFree   bool
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
