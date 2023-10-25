package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	Colony1      Colony
	ColonyMatrix [][]Room
	XMax         int
	YMax         int
	TotalPaths   string
)

func Validate(data string) bool { // Error handling later?
	isOk := true
	dataArray := strings.Split(data, "\n")
	ants, err := strconv.Atoi(dataArray[0])
	if err != nil {
		fmt.Println("Invalid data format: invalid number of ants")
		return false
	}
	Colony1.Ants = ants
	var rooms []string
	var links []string
	var testLinks bool
	start := false
	end := false
	for key, value := range dataArray[1:] {
		if strings.HasPrefix(value, "##") {
			if value == "##start" {
				if !start {
					Colony1.Start = ValidateRoom(dataArray[key+2])
					start = true
				} else {
					fmt.Println("Invalid data format: found duplicate ##start")
					return false
				}
			} else if value == "##end" {
				if !end {
					Colony1.End = ValidateRoom(dataArray[key+2])
					end = true
				} else {
					fmt.Println("Invalid data format: found duplicate ##end")
					return false
				}
			}
		} else if strings.Contains(value, " ") {
			x := strings.Split(value, " ")
			if x[0] == Colony1.End.Name || x[0] == Colony1.Start.Name {
				continue
			} else {
				rooms = append(rooms, value)
			}
		} else if strings.Contains(value, "-") {
			links = append(links, value)
		} else if value == "" {
			continue
		} else {
			fmt.Println("Invalid data format: weird format")
			return false
		}
	}
	if !start || !end {
		fmt.Println("Invalid data format: missing start or end")
		return false
	}

	for _, value := range rooms {
		Colony1.Rooms = append(Colony1.Rooms, ValidateRoom(value))
	}

	testLinks = CheckLinks(links)
	fmt.Println(Colony1)
	if !testLinks {
		
		fmt.Println("Invalid data format: invalid links")
		return false
	}
	var path []string
	path = append(path, Colony1.Start.Name)
	FindPaths(path) // Gets total paths as a string 
	totalPathsArray := strings.Split(TotalPaths, "!")    // change paths from string to [][]string
	for i:=0;i<len(totalPathsArray)-1;i++{
		totalPathsArrayNew := strings.Split(totalPathsArray[i], " ")
		Colony1.Paths = append(Colony1.Paths, totalPathsArrayNew)
	}
	Colony1.Paths = SortPaths(Colony1.Paths)
fmt.Println(Colony1.Paths)
UpdateLinks()
fmt.Println(Colony1.Rooms)
// 	matrix, err := RoomMatrix(Colony1.Rooms, 1, 2)
// 	if err != nil {
// 		fmt.Println("Invalid data format: duplicate rooms")
// 		return false
// 	}
// 	ColonyMatrix = matrix
	return isOk
}

type Colony struct {
	Ants  int
	Rooms []Room
	Start Room
	End   Room
	Paths  [][]string
}

type Room struct {
	Name     string
	X        int
	Y        int
	Links    []string
	NewLinks []FinalLinks
	Occupied bool
	Occupier int
	NumOfAnts int
}
type FinalLinks struct {
	ForwardLinks 	string
	RoomsToTheEnd   int
}
