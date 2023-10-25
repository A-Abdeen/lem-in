package lemin

func MoveAnts() {
	movedAnt := false
	for i := 0; i < len(Colony1.AntTracker); i++ {
		for j := 0; j < len(Colony1.AntTracker[i].Location.Links); j++ {
			nextRoom := Colony1.AntTracker[i].Location.Links[j]
			for k := 0; k < len(Colony1.Rooms); k++ {
				if Colony1.Rooms[k].Name == nextRoom && !Colony1.Rooms[k].Occupied {
					movedAnt = true
				} else {
					break
				}
			}
			if movedAnt {
				break
			}
		}
	}
}
