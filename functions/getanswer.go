package lemin

import (
	"log"
	"strconv"
	"strings"
)

func GetAnswer(data string) {
	dataArray := strings.Split(data, "\n")
	ants, err := strconv.Atoi(dataArray[0]) // first line in the text file should be number of ants if not return error
	if err != nil || ants <= 0 {
		log.Fatalf("Invalid data format: invalid number of ants")
	}
	Colony1.Ants = ants // Colony is a type created which well hold all the data arranged.
	var testLinks bool
	rooms, links := ValidateText(dataArray) // checks the data for error then gets the rooms and links
	for _, value := range rooms {
		Colony1.Rooms = append(Colony1.Rooms, ValidateRoom(value)) // use the func to divide the data for each room
	}
	Matrix = RoomMatrix(Colony1.Rooms, XMax, YMax)
	//PrintMatrix() 												  // Display static grid with plotted rooms
	testLinks = CheckLinks(links) // check the links for error and add the links to each room
	if !testLinks {
		log.Fatalf("Invalid data format: invalid links")
	}
	var path []string
	path = append(path, Colony1.Start.Name)  // find paths starting with start room
	FindPaths(path)                          // well add all the paths from start to end to Colony1
	StringToArray()                          // change paths from string to [][]string
	Colony1.Paths = SortPaths(Colony1.Paths) // sort paths from shortest to longest
	UpdateLinks()                            // add links only if they are moving forward and arranged from the link with the shortest path to the longest
	for i := 1; i <= Colony1.Ants; i++ {     // Create an ant tracker that well have all the ants along with the room there in
		Colony1.AntTracker = append(Colony1.AntTracker, Ant{i, Colony1.Start})
	}
	MoveAnts() // function well give the final answer using the data in the global variable
}
