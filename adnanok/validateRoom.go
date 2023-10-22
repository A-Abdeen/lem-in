package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateRoom(rooms []string) (int, int) {
	xMax := 0
	yMin := 0

	for _, value := range rooms {
		current := strings.Split(value, " ")
		if len(current) != 3 {
			fmt.Println("Invalid data format: invalid room format")
			os.Exit(0)
		}
		if current[0] == "" || current[1] == "" || current[2] == "" {
			fmt.Println("Invalid data format: invalid room format")
			os.Exit(0)
		}
		xCord, err := strconv.Atoi(current[1])
		if err != nil {
			fmt.Println("Invalid data format: invalid room format")
			os.Exit(0)
		}
		yCord, err := strconv.Atoi(current[2])
		if err != nil {
			fmt.Println("Invalid data format: invalid room format")
			os.Exit(0)
		}
		if xCord > xMax {
			xMax = xCord
		}
		if yCord < yMin {
			yMin = yCord
		}

		Colony1.Rooms = append(Colony1.Rooms, Room{Name: current[0], X: xCord, Y: yCord})
	}
	return xMax, yMin
}
