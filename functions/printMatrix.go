package lemin

import "fmt"

func PrintMatrix() {
	for i := 0; i < YMax+2; i++ {
		for j := 0; j < XMax+2; j++ {
			if Matrix[j][i] == "" {
				Matrix[j][i] = "[  ]\t"
			}
			fmt.Print(Matrix[j][i])
		}
		fmt.Println()
	}
	fmt.Println()

}
