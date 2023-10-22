package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateRoom(rooms []string) {
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
		
		Colony1.Rooms = append(Colony1.Rooms, Room{Name: current[0], X: xCord, Y: yCord})
	}
}
