package lemin

import "log"

func RoomMatrix(rooms []Room, x, y int) [][]string {
	x += 2
	y += 2
	matrix := make([][]string, x)
	for i := range matrix {
		matrix[i] = make([]string, y)
	}
	matrix[Colony1.Start.X][Colony1.Start.Y] = Colony1.Start.Name
	matrix[Colony1.End.X][Colony1.End.Y] = Colony1.End.Name
	for i := range rooms {
		if matrix[rooms[i].X][rooms[i].Y] != "" {
			log.Fatalf("Invalid data format: duplicate coordinates")
		} else {
			matrix[rooms[i].X][rooms[i].Y] = rooms[i].Name
		}
	}
	return matrix
}
