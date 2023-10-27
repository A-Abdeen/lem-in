package lemin

import (
	"fmt"
	"strconv"
	// "os"
)

func MoveAnts() {
	var numbofMoves int
	var bestWayToStart int
	var finalcount int
	// for st:=(len(Colony1.Start.NewLinks));st>0;st--{
	// 	if Colony1.Ants%st == 0 {
	// 		bestWayToStart = st
	// 		break
	// 	}
	// }
	for i:=len(Colony1.Start.NewLinks);i>0;i--{
		count:=0
		numberOfAnts := Colony1.Ants
		for st:=i;st>0;{
			// fmt.Println(count)
			numberOfAnts= numberOfAnts-st
			if numberOfAnts == 0 {
				count++
				break
			} else if numberOfAnts>0{
				count++
				continue
			} else if numberOfAnts<0{
		
				numberOfAnts = numberOfAnts+st
				st--
			}
		}
		fmt.Println(count)
		if count <= finalcount || finalcount == 0{
			finalcount = count
			bestWayToStart = i
			// fmt.Println(bestWayToStart)
		}
	}
	// os.Exit(0)
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
				for k := 0; k < len(Colony1.Rooms); k++ { // check if the room that the ant links to is not Occupied
					
						if nextRoomLink == Colony1.End.Name && !movedAntToEnd{
						if currentRoom.Name == Colony1.Start.Name{
						movedAntToEnd = true}
						movedAnt = true
						fmt.Print("L" + strconv.Itoa(Colony1.AntTracker[i].Number) + "-" + Colony1.End.Name + "\t")
						Colony1.End.NumOfAnts++
						Colony1.AntTracker[i].Location = Colony1.End
						break
					}
					
					if Colony1.Rooms[k].Name == nextRoomLink{
						if Colony1.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd ==1 && ((Colony1.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd - Colony1.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd)>(Colony1.Start.NumOfAnts)){
						} else if (Colony1.AntTracker[i].Location.NewLinks[j].RoomsToTheEnd - Colony1.AntTracker[i].Location.NewLinks[0].RoomsToTheEnd)>(Colony1.Start.NumOfAnts+2){
						} else if !Colony1.Rooms[k].Occupied {
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
					if currentRoom.Name == Colony1.Start.Name {
						Colony1.Start.NumOfAnts--
					}
					for m := 0; m < len(Colony1.Rooms); m++ {
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
