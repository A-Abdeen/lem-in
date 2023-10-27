package lemin

import "strings"

func StringToArray() {
	totalPathsArray := strings.Split(TotalPaths, "!") // change paths from string to [][]string
	for i := 0; i < len(totalPathsArray)-1; i++ {
		totalPathsArrayNew := strings.Split(totalPathsArray[i], " ")
		Colony1.Paths = append(Colony1.Paths, totalPathsArrayNew)
	}
}
