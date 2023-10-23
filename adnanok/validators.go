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
	var structStart string
	var structEnd string
	var testLinks bool
	start := false
	end := false
	for key, value := range dataArray[1:] {
		if strings.HasPrefix(value, "##") {
			if value == "##start" {
				if !start {
					Colony1.Start = ValidateRoom(dataArray[key+2])
					fmt.Println("START ", structStart)
					start = true
				} else {
					fmt.Println("Invalid data format: found duplicate ##start")
					return false
				}
			} else if value == "##end" {
				if !end {
					Colony1.End = ValidateRoom(dataArray[key+2])
					fmt.Println("END ", structEnd)
					end = true
				} else {
					fmt.Println("Invalid data format: found duplicate ##end")
					return false
				}
			}
		} else if strings.Contains(value, " ") {
			if value == structEnd || value == structStart {
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
	fmt.Println(ants)

	for _, value := range rooms {
		Colony1.Rooms = append(Colony1.Rooms, ValidateRoom(value))
	}
	fmt.Println(rooms)

	testLinks = CheckLinks(Colony1.Rooms, links)
	if !testLinks {
		fmt.Println("Invalid data format: invalid links")
		return false
	}
	fmt.Println(links)

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
}

type Room struct {
	Name     string
	X        int
	Y        int
	Links    []string
	Occupied bool
}
