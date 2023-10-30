package lemin

func FindWayToStart() int {
	var finalcount int
	var bestWayToStart int
	for i := len(Colony1.Start.NewLinks); i > 0; i-- {
		count := 0
		numberOfAnts := Colony1.Ants
		for st := i; st > 0; {
			numberOfAnts = numberOfAnts - st
			if numberOfAnts == 0 {
				count++
				break
			} else if numberOfAnts > 0 {
				count++
				continue
			} else if numberOfAnts < 0 {
				numberOfAnts = numberOfAnts + st
				st--
			}
		}
		if count <= finalcount || finalcount == 0 {
			finalcount = count
			bestWayToStart = i
		}
	}
	return bestWayToStart
}
