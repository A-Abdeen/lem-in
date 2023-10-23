package lemin

import (
	"fmt"
	"strings"
)

func CheckLinks(links []string) bool {
	for _, link := range links {
		// Check if room exists
		availableRoom1 := false
		availableRoom2 := false
		room1, room2, err := strings.Cut(link, "-")
		if !err || room1 == room2 {
			fmt.Println("Invalid data format: wrong link name1")
			return false
		}
		// Find Room1 of above link
		for i := 0; i < len(Colony1.Rooms); i++ {
			if room1 == Colony1.Rooms[i].Name {
				// Validate existing links
				for _, testinglinks := range Colony1.Rooms[i].Links {
					if testinglinks == room2 {
						fmt.Println("Invalid data format: wrong link name2")
						return false
					}
				}
				// Add Room2 to list of Room1 links
				Colony1.Rooms[i].Links = append(Colony1.Rooms[i].Links, room2)
				availableRoom1 = true
			} else if room1 == Colony1.Start.Name {
				// Validate existing links
				for _, testinglinks := range Colony1.Start.Links {
					if testinglinks == room2 {
						fmt.Println("Invalid data format: wrong link name3")
						return false
					}
				}
				// Add Room2 to list of Room1 links
				Colony1.Start.Links = append(Colony1.Start.Links, room2)
				availableRoom1 = true

			} else if room1 == Colony1.End.Name {
				// Validate existing links
				for _, testinglinks := range Colony1.End.Links {
					if testinglinks == room2 {
						fmt.Println("Invalid data format: wrong link name4")
						return false
					}
				}
				// Add Room2 to list of Room1 links
				Colony1.End.Links = append(Colony1.End.Links, room2)
				availableRoom1 = true
			} else if room2 == Colony1.Rooms[i].Name { // TO LOOK AT THE SECOND ROOM LINKS
				// Validate existing links
				for _, testinglinks := range Colony1.Rooms[i].Links {
					if testinglinks == room1 {
						fmt.Println("Invalid data format: wrong link name5")
						return false
					}
				}
				// Add Room1 to list of Room2 links
				Colony1.Rooms[i].Links = append(Colony1.Rooms[i].Links, room1)
				availableRoom2 = true
			} else if room2 == Colony1.Start.Name { // TO LOOK AT THE SECOND ROOM LINKS
				// Validate existing links
				for _, testinglinks := range Colony1.Start.Links {
					if testinglinks == room1 {
						fmt.Println("Invalid data format: wrong link name6")
						return false
					}
				}
				// Add Room1 to list of Room2 links
				Colony1.Start.Links = append(Colony1.Start.Links, room1)
				availableRoom2 = true
			} else if room2 == Colony1.End.Name { // TO LOOK AT THE SECOND ROOM LINKS
				// Validate existing links
				for _, testinglinks := range Colony1.End.Links {
					if testinglinks == room1 {
						fmt.Println("Invalid data format: wrong link name7")
						return false
					}
				}
				// Add Room1 to list of Room2 links
				Colony1.End.Links = append(Colony1.End.Links, room1)
				availableRoom2 = true
			} else {
				continue
			}
			if !availableRoom1 || !availableRoom2 {
				fmt.Println("ASAP: ", availableRoom1, availableRoom2)
				return false
			}
		}
	}
	return true
}
