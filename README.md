type Parkinglot struct{
    id 
    address 
    entrance 
    exit
    parkingfloor
}

type parkingfloor struct{
    id
    floor no
    slots
    display
}

type entrance struct{
    id
}

type exit struct{
    id
}

type slots struct{
    id
    slottype
    isfull

}

type display struct{
    avaiableslots
}

type slottype struct{
    car iota
    byke
}

type vehicle struct{
    registration
    vehicletype
    color
}# parking_lot_golang
