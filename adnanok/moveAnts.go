package lemin

import (
	"fmt"
	"strconv"
)

func MoveAnts() {
	var numbofMoves int
	bestWayToStart := FindWayToStart() 
	for h := 0; h >= 0; h++ {
		movedAntToEnd := false
		if Colony1.Ants == Colony1.End.NumOfAnts {
			break
		}
		for i := 0; i < len(Colony1.AntTracker); i++ { //loop over ants one at a time
			for j := 0; j < len(Colony1.AntTracker[i].Location.NewLinks); j++ { //loop over the links of the room that the ant is in
				currentRoom := Colony1.AntTracker[i].Location
				nextRoomLink := Colony1.AntTracker[i].Location.NewLinks[j].ForwardLinks
				movedAnt := false
				if  (currentRoom.Name == Colony1.Start.Name && j>=bestWayToStart){
					continue
				}
				for k := 0; k < len(Colony1.Rooms); k++ { // go over each room to check the links
						if nextRoomLink == Colony1.End.Name && !movedAntToEnd{ // check if the end the room is linked 
						if currentRoom.Name == Colony1.Start.Name{
						movedAntToEnd = true
					}
						movedAnt = true
						fmt.Print("L" + strconv.Itoa(Colony1.AntTracker[i].Number) + "-" + Colony1.End.Name + "\t")
						Colony1.End.NumOfAnts++
						Colony1.AntTracker[i].Location = Colony1.End
						break
					}
					
					if Colony1.Rooms[k].Name == nextRoomLink{ // if the link matches the name check if the difference between the link and the best link which has the least rooms to the end
						if Colony1.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd ==1 && ((Colony1.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd - Colony1.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd)>(Colony1.Start.NumOfAnts)){
						} else if (Colony1.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd - Colony1.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd)>(Colony1.Start.NumOfAnts+2){
						} else if !Colony1.Rooms[k].Occupied {// check if the room that the ant links to is not Occupied
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
				if movedAnt { // counter for the number of ants in the start room
					if currentRoom.Name == Colony1.Start.Name {
						Colony1.Start.NumOfAnts--
					}
					for m := 0; m < len(Colony1.Rooms); m++ { // change the boolean of the room the ant moved to. 
						if currentRoom.Name == Colony1.Rooms[m].Name {
							Colony1.Rooms[m].Occupied = false
							break
						}
					}
					break
				}
			}
		}
		fmt.Println()
		numbofMoves++
	}
	fmt.Println(numbofMoves)
}
