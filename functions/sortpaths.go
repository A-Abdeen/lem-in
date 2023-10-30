package lemin

func SortPaths(paths [][]string) [][]string {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) < len(paths[j]) {
				continue
			} else {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
	return paths
}
