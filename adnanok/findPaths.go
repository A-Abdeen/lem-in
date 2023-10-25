package lemin

func FindPaths(path []string) bool {
	var selectedRoom Room
	which := path[len(path)-1]
	if which == Colony1.End.Name { // if the final string in the array path is the END add the data to the string and return
		for i:=0;i<len(path);i++{
		if i != len(path)-1 {
			TotalPaths = TotalPaths + path[i]+ " "
			} else {
				TotalPaths = TotalPaths + path[i]
			} 
		}
		TotalPaths = TotalPaths + "!"
		return true
	}
	for i := 0; i < len(Colony1.Rooms); i++ { // find the last string and add the data of the Room which matches the name
		if which == Colony1.Start.Name {
			selectedRoom = Colony1.Start
			break
		}
		if which == Colony1.Rooms[i].Name {
			selectedRoom = Colony1.Rooms[i]
			break
		}
	}
	for i := 0; i < len(selectedRoom.Links); i++ { // check all the links for the selected link
		beenThere := false
		newpath := path
		for j := 0; j < len(path); j++ {  // if the link matches a previous room in the array break
			if selectedRoom.Links[i] == path[j] {
				beenThere = true
				break
			}
		}
		if beenThere {
			continue
		}
		newpath = append(newpath, selectedRoom.Links[i]) // append the new room and enter the function again
		FindPaths(newpath)
	}
	return false
}
