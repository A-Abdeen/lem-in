package lemin

func RoomMatrix(rooms []Room, x, y int) ([][]int, error) {
	matrix := make([][]int, y)
	for i := range matrix {
		matrix[i] = make([]int, x)
	}
	// Loop over rooms and assign to matrix
	return matrix, nil
}
