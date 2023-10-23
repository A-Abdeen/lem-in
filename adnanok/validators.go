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
	if !testLinks {
		fmt.Println(Colony1)
		fmt.Println("Invalid data format: invalid links")
		return false
	}

	matrix, err := RoomMatrix(Colony1.Rooms, 1, 2)
	if err != nil {
		fmt.Println("Invalid data format: duplicate rooms")
		return false
	}
	ColonyMatrix = matrix
	return isOk
}

type Colony struct {
	Ants  int
	Rooms []Room
	Start Room
	End   Room
	Path  [][]string
}

type Room struct {
	Name     string
	X        int
	Y        int
	Links    []string
	Occupied bool
	Occupier int
}
