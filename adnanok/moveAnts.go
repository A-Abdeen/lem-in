package lemin

import (
	"fmt"
	"strconv"
)

func MoveAnts() {
	for h := 0; h >= 0; h++ {
		if Colony1.Ants == Colony1.End.NumOfAnts {
			break
		}
		for i := 0; i < len(Colony1.AntTracker); i++ { //loop over ants one at a time
			for j := 0; j < len(Colony1.AntTracker[i].Location.NewLinks); j++ { //loop over the links of the room that the ant is in
				nextRoom := Colony1.AntTracker[i].Location
				nextRoomLink := Colony1.AntTracker[i].Location.NewLinks[j].ForwardLinks
				movedAnt := false
				for k := 0; k < len(Colony1.Rooms); k++ { // check if the room that the ant links to is not Occupied
					if nextRoomLink == Colony1.End.Name {
						movedAnt = true
						fmt.Print("L" + strconv.Itoa(Colony1.AntTracker[i].Number) + "-" + Colony1.End.Name + "\t")
						Colony1.End.NumOfAnts++
						Colony1.AntTracker[i].Location = Colony1.End
						break
					}
					if Colony1.Rooms[k].Name == nextRoomLink {
						if !Colony1.Rooms[k].Occupied {
							Colony1.Rooms[k].Occupied = true
							fmt.Print("L" + strconv.Itoa(Colony1.AntTracker[i].Number) + "-" + Colony1.Rooms[k].Name + "\t")
							Colony1.AntTracker[i].Location = Colony1.Rooms[k]
							movedAnt = true
							break
						} else {
							break
						}
					}
				}
				if movedAnt {
					for m := 0; m < len(Colony1.Rooms); m++ {
						if nextRoom.Name == Colony1.Rooms[m].Name {
							Colony1.Rooms[m].Occupied = false
							break
						}
					}
					break
				}
			}
		}
		fmt.Println()
	}
}
