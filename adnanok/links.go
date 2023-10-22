package lemin

import (
	"fmt"
	"strings"
)

func CheckLinks(Rooms []Room, links []string) bool {
	for _, link := range links{
	availableRoom1 := false
	availableRoom2 := false
	room1, room2, err := strings.Cut(link, "-")
	if !err|| room1 == room2 {
		fmt.Println("Invalid data format: wrong link name")
	}
	for i := 0; i < len(Rooms); i++ {
		if room1 == Colony1.Rooms[i].Name {
			for _, testinglinks := range Colony1.Rooms[i].Links {
				if testinglinks == room2 {
					fmt.Println("Invalid data format: wrong link name")
					return false
				}
			}
			Colony1.Rooms[i].Links = append(Colony1.Rooms[i].Links, room2)
			availableRoom1 = true
		}
		if room2 == Colony1.Rooms[i].Name {
			for _, testinglinks := range Colony1.Rooms[i].Links {
				if testinglinks == room1 {
					fmt.Println("Invalid data format: wrong link name")
					return false
				}
			}
			Colony1.Rooms[i].Links = append(Colony1.Rooms[i].Links, room1)
			availableRoom2 = true
		}
	}
	if !availableRoom1 == !availableRoom2 {
		return false
	} 
}
	return true
}
