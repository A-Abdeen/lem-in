package lemin

func UpdateLinks() {
	var tempFinalLinks FinalLinks
	for i := 0; i < len(Colony1.Paths); i++ { // go over the sorted paths from shortest to longest
		path := Colony1.Paths[i]
		for m := 0; m < len(path); m++ { // go over the path room by room 
			linkThere := false
			if Colony1.Start.Name == path[m] { // if the name matches the room check its links
				for k := 0; k < len(Colony1.Start.NewLinks); k++ { 
					if (m + 1) < len(path) {
						if (Colony1.Start.NewLinks[k].ForwardLinks) == path[m+1] {
							linkThere = true
							break
						}
					}
				}
				if (m+1) < len(path) && !linkThere { // if the next name on the path doesnot match one of the newlinks append it 
					tempFinalLinks.ForwardLinks = path[m+1]
					tempFinalLinks.RoomsToTheEnd = (len(path) - m - 1) // add the number of rooms to the end to the newlinks struct
					tempFinalLinks.lengthOfPath = len(path) // add the total length of the path to the struct
					Colony1.Start.NewLinks = append(Colony1.Start.NewLinks, tempFinalLinks)
				}
			}
			for j := 0; j < len(Colony1.Rooms); j++ { // to check all the rooms
				linkThere := false
				if Colony1.Rooms[j].Name == path[m] { 
					for k := 0; k < len(Colony1.Rooms[j].NewLinks); k++ {
						if (m + 1) < len(path) {
							if (Colony1.Rooms[j].NewLinks[k].ForwardLinks) == path[m+1] {
								linkThere = true
								break
							}
						}
					}

					if (m+1) < len(path) && !linkThere {
						tempFinalLinks.ForwardLinks = path[m+1]
						tempFinalLinks.RoomsToTheEnd = (len(path) - m - 1)
						tempFinalLinks.lengthOfPath = len(path)
						Colony1.Rooms[j].NewLinks = append(Colony1.Rooms[j].NewLinks, tempFinalLinks)
					}
				}
			}
		}
	}
}
