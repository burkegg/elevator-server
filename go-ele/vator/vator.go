package vator

import (
	"crypto/sha1"
	"fmt"
)

// FloorDesc - basic model for Floor identity
type FloorDesc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CarDesc - basic model for Car identity
type CarDesc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Vator - modeled elevator
type Vator struct {
	floors       []FloorDesc
	lookupFloors map[string]FloorDesc
	lookupCars   map[string]CarDesc
	carAtFloor   map[string]string
}

// NewVator - instantiate a new elevator model
func NewVator(floors []string, carCt int8) (*Vator, error) {
	order := []FloorDesc{}
	lookFloor := map[string]FloorDesc{}
	for _, fname := range floors {
		h := sha1.New()
		h.Write([]byte("Floor-"))
		h.Write([]byte(fname))
		fid := fmt.Sprintf("%x", h.Sum(nil))
		fDesc := FloorDesc{ID: fid, Name: fname}
		order = append(order, fDesc)
		lookFloor[fid] = fDesc
	}
	firstFloor := order[0].ID

	lookCar := map[string]CarDesc{}
	atFloor := map[string]string{}
	for ix := int8(0); ix < carCt; ix++ {
		cname := fmt.Sprintf("Car-%d", ix)
		h := sha1.New()
		h.Write([]byte(cname))
		cid := fmt.Sprintf("%x", h.Sum(nil))
		cDesc := CarDesc{ID: cid, Name: cname}
		lookCar[cid] = cDesc
		atFloor[cid] = firstFloor
	}

	v := &Vator{floors: order, lookupFloors: lookFloor, lookupCars: lookCar, carAtFloor: atFloor}
	return v, nil
}

// Cars - get a list of all the cars
func (v *Vator) Cars() []CarDesc {
	results := make([]CarDesc, 0, len(v.lookupCars))
	for _, cDesc := range v.lookupCars {
		results = append(results, cDesc)
	}
	return results
}

// Floors - get a list of all the floors
func (v *Vator) Floors() []FloorDesc {
	return v.floors
}

// Current - what floor is the car on
func (v *Vator) Current(carID string) FloorDesc {
	fid := v.carAtFloor[carID]
	return v.lookupFloors[fid]
}

// GetNearestCar - which car is located on the closest floor to the given floorID
func (v *Vator) GetNearestCar(floorID string) (carID string) {
	// To be copmleted by candidate
	return ""
}

// CallCar - move nearest car to the specified floor
func (v *Vator) CallCar(floorID string) (err error) {
	// To be completed by candidate
	return nil
}
