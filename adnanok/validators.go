package lemin

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Validate(data string) {
	dataArray := strings.Split(data, "\n")
	ants, err := strconv.Atoi(dataArray[0])
	if err != nil {
		log.Fatalf("Invalid data format: invalid number of ants")
	}
	Colony1.Ants = ants
	var rooms []string
	var links []string
	var testLinks bool
	start := false
	end := false
	for key, value := range dataArray[1:] {
		if strings.HasPrefix(value, "#") {
			if value == "##start" {
				if !start {
					Colony1.Start = ValidateRoom(dataArray[key+2])
					Colony1.Start.NumOfAnts = Colony1.Ants
					Colony1.Start.Occupied = true
					start = true
				} else {
					log.Fatalf("Invalid data format: found duplicate ##start")
				}
			} else if value == "##end" {
				if !end {
					Colony1.End = ValidateRoom(dataArray[key+2])
					end = true
				} else {
					log.Fatalf("Invalid data format: found duplicate ##end")
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
			log.Fatalf("Invalid data format: weird format")
		}
	}
	if !start || !end {
		log.Fatalf("Invalid data format: missing start or end")
	}

	for _, value := range rooms {
		Colony1.Rooms = append(Colony1.Rooms, ValidateRoom(value))
	}
	for i := 1; i <= Colony1.Ants; i++ {
		Colony1.AntTracker = append(Colony1.AntTracker, Ant{i, Colony1.Start})
	}
	Matrix = RoomMatrix(Colony1.Rooms, XMax, YMax)
	PrintMatrix() // Display static grid with plotted rooms
	testLinks = CheckLinks(links)
	if !testLinks {
		log.Fatalf("Invalid data format: invalid links")
	}
	var path []string
	path = append(path, Colony1.Start.Name)
	FindPaths(path)                                   // Gets total paths as a string
	totalPathsArray := strings.Split(TotalPaths, "!") // change paths from string to [][]string
	for i := 0; i < len(totalPathsArray)-1; i++ {
		totalPathsArrayNew := strings.Split(totalPathsArray[i], " ")
		Colony1.Paths = append(Colony1.Paths, totalPathsArrayNew)
	}
	Colony1.Paths = SortPaths(Colony1.Paths)
	fmt.Println((Colony1.Paths))
	UpdateLinks()
	fmt.Println(Colony1.Rooms)
	for i := 1; i <= Colony1.Ants; i++ {
		Colony1.AntTracker = append(Colony1.AntTracker, Ant{i, Colony1.Start})
	}
	MoveAnts()
 }

func PrintMatrix() {
	for i := 0; i < YMax+2; i++ {
		for j := 0; j < XMax+2; j++ {
			if Matrix[j][i] == "" {
				Matrix[j][i] = "[  ]\t"
			}
			fmt.Print(Matrix[j][i])
		}
		fmt.Println()
	}
	fmt.Println()

}
