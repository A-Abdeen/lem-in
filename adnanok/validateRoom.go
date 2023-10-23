package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateRoom(value string) Room {
	current := strings.Split(value, " ")
	if len(current) != 3 {
		fmt.Println("Invalid data format: invalid room format")
		os.Exit(0)
	}
	if current[0] == "" || current[1] == "" || current[2] == "" {
		fmt.Println("Invalid data format: invalid room format")
		os.Exit(0)
	}
	if strings.HasPrefix(current[0], "L") || strings.HasPrefix(current[0], "#") {
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
	if xCord > XMax {
		XMax = xCord
	}
	if yCord < YMax {
		YMax = yCord
	}
	return Room{Name: current[0], X: xCord, Y: yCord}
}
