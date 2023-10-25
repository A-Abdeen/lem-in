package main

import (
	"encoding/json"
	"fmt"
	lemin "lemin/adnanok"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage case requires txt file")
	}
	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var dataModified []byte
	for i := 0; i < len(data); i++ {
		if data[i] != 13 {
			dataModified = append(dataModified, data[i])
		}
	}
	dataStr := string(dataModified)
	lemin.Validate(dataStr)
	PrintJSON(lemin.Colony1)
}
func PrintJSON(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	fmt.Println(string(bytes))
}
