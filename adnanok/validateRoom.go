package lemin

import (
	"log"
	"strconv"
	"strings"
)

func ValidateRoom(value string) Room {
	current := strings.Split(value, " ")
	if len(current) != 3 {
		log.Fatalf("Invalid data format: invalid room format")
	}
	if current[0] == "" || current[1] == "" || current[2] == "" {
		log.Fatalf("Invalid data format: invalid room format")
	}
	if strings.HasPrefix(current[0], "L") || strings.HasPrefix(current[0], "#") {
		log.Fatalf("Invalid data format: invalid room format")
	}
	xCord, err := strconv.Atoi(current[1])
	if err != nil {
		log.Fatalf("Invalid data format: invalid room format")
	}
	yCord, err := strconv.Atoi(current[2])
	if err != nil {
		log.Fatalf("Invalid data format: invalid room format")
	}
	if xCord > XMax {
		XMax = xCord
	}
	if yCord > YMax {
		YMax = yCord
	}
	return Room{Name: current[0], X: xCord, Y: yCord}
}
