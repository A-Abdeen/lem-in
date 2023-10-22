package lemin

import (
	"fmt"
	"strings"
)

func Links(Rooms []Room, link string) ([]Room, bool) {
	availableRoom1 := false
	availableRoom2 := false
	room1, room2, err := strings.Cut(link, "-")
	if !err|| room1 == room2 {
		fmt.Println("Invalid data format: wrong link name")
	}
	for i := 0; i < len(Rooms); i++ {
		if room1 == Rooms[i].Name {
			Rooms[i].Links = append(Rooms[i].Links, room2)
			availableRoom1 = true
		}
		if room2 == Rooms[i].Name {
			Rooms[i].Links = append(Rooms[i].Links, room1)
			availableRoom2 = true
		}
	}
	if !availableRoom1 == !availableRoom2 {
		return nil, false
	} 
	return Rooms, true
}
