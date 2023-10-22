package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(data string) bool { // Error handling later?
	var colony Colony
	isOk := true
	dataArray := strings.Split(data, "\n")
	ants, err := strconv.Atoi(dataArray[0])
	if err != nil {
		fmt.Println("Invalid data format: invalid number of ants")
		return false
	}
	colony.Ants = ants
	var rooms []string
	var links []string
	var structStart string
	var structEnd string
	start := false // implement struct func
	end := false
	for key, value := range dataArray[1:] {
		if strings.HasPrefix(value, "##") {
			if value == "##start" {
				if !start {
					structStart = dataArray[key+2]
					fmt.Println("START ", structStart)
					start = true
				} else {
					fmt.Println("Invalid data format: found duplicate ##start")
					return false
				}
			} else if value == "##end" {
				if !end {
					structEnd = dataArray[key+2]
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

	// Call function that validates rooms
	fmt.Println(rooms)

	// Call function that validates links
	fmt.Println(links)

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
	Links    []Link
	Occupied bool
}

type Link struct {
	RoomA Room
	RoomB Room
}
