package lemin

func RoomMatrix(rooms []Room, x, y int) ([][]Room, error) {
	matrix := make([][]Room, y)
	for i := range matrix {
		matrix[i] = make([]Room, x)
	}
	// Loop over rooms and assign to matrix
	for i := range rooms {
		matrix[rooms[i].X][rooms[i].Y] = rooms[i]
	}
	return matrix, nil
}
