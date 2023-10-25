package lemin

// import "fmt"

func RoomMatrix(rooms []Room, x, y int) [][]Room {
	matrix := make([][]Room, x)
	for i := range matrix {
		matrix[i] = make([]Room, y)
	}
	// fmt.Println(matrix)
	// Loop over rooms and assign to matrix
	// for i := range rooms {
	// 	matrix[rooms[i].X][rooms[i].Y] = rooms[i]
	// }
	return matrix
}
