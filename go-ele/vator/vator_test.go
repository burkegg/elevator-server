package vator

import "testing"

// helper to get floor id from name
func getFloorIdByName(vator *Vator, name string) (floorID string) {
	floors := vator.Floors()
	for _, floor := range floors {
		if floor.Name == name {
			return floor.ID
		}
	}
	return ""
}

// helper to get floor object from floor ID
func getFloorByFloorID(vator *Vator, floorID string) (floor *FloorDesc) {
	floors := vator.Floors()
	for _, floor := range floors {
		if floor.ID == floorID {
			return &floor
		}
	}
	return nil
}

func TestInstantiate(t *testing.T) {
	vator, err := NewVator([]string{"B3", "B2", "B1", "Lobby", "F1", "F2", "F3"}, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	t.Logf("Order: %v\n", vator.floors)
	t.Logf("Second: %v\n", vator.Floors())
}

func TestGetNearestCar(t *testing.T) {
	vator, err := NewVator([]string{"B3", "B2", "B1", "Lobby", "F1", "F2", "F3", "F4"}, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	cars := vator.Cars()
	carOne, carTwo, carThree := cars[0], cars[1], cars[2]
	topFloorID := getFloorIdByName(vator, "F4")
	lobbyID := getFloorIdByName(vator, "Lobby")
	vator.carAtFloor[carOne.ID] = topFloorID
	vator.carAtFloor[carTwo.ID] = lobbyID
	b2FloorID := getFloorIdByName(vator, "B2")
	closestToB2 := vator.GetNearestCar(b2FloorID)
	if closestToB2 != carThree.ID {
		t.Errorf("Unexpected values for %s, %s", closestToB2, carThree.ID)
	}
	f1FloorID := getFloorIdByName(vator, "F1")
	closestToF1 := vator.GetNearestCar(f1FloorID)
	if closestToF1 != carTwo.ID {
		t.Errorf("Unexpected values for %s, %s", closestToF1, carTwo.ID)
	}
	f3FloorID := getFloorIdByName(vator, "F3")
	closestToF3 := vator.GetNearestCar(f3FloorID)
	if closestToF3 != carOne.ID {
		t.Errorf("Unexpected values for %s, %s", closestToF3, carOne.ID)
	}
}

func TestCallToFloor(t *testing.T) {
	vator, err := NewVator([]string{"B4", "B3", "B2", "B1", "Lobby", "F1", "F2", "F3", "F4", "F5"}, 3)
	cars := vator.Cars()
	carOne, carTwo, carThree := cars[0], cars[1], cars[2]
	topFloorID := getFloorIdByName(vator, "F5")
	lobbyID := getFloorIdByName(vator, "Lobby")
	lobbyFloor := getFloorByFloorID(vator, lobbyID)
	vator.carAtFloor[carOne.ID] = topFloorID
	vator.carAtFloor[carTwo.ID] = lobbyID
	b3FloorID := getFloorIdByName(vator, "B3")
	floorB3 := getFloorByFloorID(vator, b3FloorID)
	err = vator.CallCar(b3FloorID)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	if vator.Current(carThree.ID) != *floorB3 {
		t.Errorf("Unexpected current floor for car %s", carThree.ID)
	}
	f4FloorID := getFloorIdByName(vator, "F4")
	floorF4 := getFloorByFloorID(vator, f4FloorID)
	err = vator.CallCar(f4FloorID)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
	if vator.Current(carOne.ID) != *floorF4 {
		t.Errorf("Unexpected current floor for car %s", carOne.ID)
	}
	// Check that the lobby car hasn't moved
	if vator.Current(carTwo.ID) != *lobbyFloor {
		t.Errorf("Unexpected current floor for car %s", carTwo.ID)
	}
}
