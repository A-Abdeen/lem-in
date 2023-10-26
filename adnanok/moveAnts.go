package lemin

import (
	// "fmt"
	"strconv"
)

func MoveAnts(colony2 Colony, numbOfTurns int)(Colony, string) {
	var finalAnswer string
		movedAntToEnd := false
		for NumbOfTurns:=0;NumbOfTurns<numbOfTurns;NumbOfTurns++{
		for i := 0; i < len(colony2.AntTracker); i++ { //loop over ants one at a time
		
			for j := 0; j < len(colony2.AntTracker[i].Location.NewLinks); j++ { //loop over the links of the room that the ant is in
				currentRoom := colony2.AntTracker[i].Location
				nextRoomLink := colony2.AntTracker[i].Location.NewLinks[j].ForwardLinks
				movedAnt := false
				for k := 0; k < len(colony2.Rooms); k++ { // check if the room that the ant links to is not Occupied
					if nextRoomLink == colony2.End.Name && !movedAntToEnd{
						if currentRoom.Name == colony2.Start.Name {
						movedAntToEnd = true}
						movedAnt = true
						finalAnswer = finalAnswer + "L" + strconv.Itoa(colony2.AntTracker[i].Number) + "-" + colony2.End.Name + "\t"
						
						colony2.End.NumOfAnts++
						colony2.AntTracker[i].Location = colony2.End
						break
					}
					if colony2.Rooms[k].Name == nextRoomLink {
						// fmt.Println(colony2.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd, colony2.AntTracker[i].Location.NewLinks[0].lengthOfPath, colony2.Start.NumOfAnts, colony2.End.NumOfAnts)
						if ((colony2.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd - colony2.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd)>(colony2.Start.NumOfAnts+2)) || (colony2.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd > numbOfTurns-2){
							} else if !colony2.Rooms[k].Occupied {					
							colony2.Rooms[k].Occupied = true
							finalAnswer = finalAnswer + ("L" + strconv.Itoa(colony2.AntTracker[i].Number) + "-" + colony2.Rooms[k].Name + "\t")
							colony2.AntTracker[i].Location = colony2.Rooms[k]
							movedAnt = true
							break
						} else {
							break
						}
					}
				}
				if movedAnt {
					if currentRoom.Name == colony2.Start.Name {
						colony2.Start.NumOfAnts--
					}
					for m := 0; m < len(colony2.Rooms); m++ {
						if currentRoom.Name == colony2.Rooms[m].Name {
							colony2.Rooms[m].Occupied = false
							break
						}
					}
					break
				}
			}
		}
		// fmt.Println(1)
	finalAnswer = finalAnswer + "\n"
	// fmt.Println(ColonyTemp)
	}
	
	return colony2, finalAnswer
}
