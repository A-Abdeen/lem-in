package lemin

func UpdateLinks() {
	var tempFinalLinks FinalLinks
	
	for i := 0; i < len(Colony1.Paths); i++ {
		path := Colony1.Paths[i]
		for m := 0; m < len(path); m++ {
			linkThere := false
			if Colony1.Start.Name == path[m] {
				for k := 0; k < len(Colony1.Start.NewLinks); k++ {
					if (m + 1) < len(path) {
						if (Colony1.Start.NewLinks[k].ForwardLinks) == path[m+1] {
							linkThere = true
							}
					}
				}
				if (m+1) < len(path) && !linkThere {
					tempFinalLinks.ForwardLinks = path[m+1]
					tempFinalLinks.RoomsToTheEnd = (len(path) - m - 1)
					tempFinalLinks.lengthOfPath = len(path)
					Colony1.Start.NewLinks = append(Colony1.Start.NewLinks, tempFinalLinks)
				}
			}
			for j := 0; j < len(Colony1.Rooms); j++ {
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
