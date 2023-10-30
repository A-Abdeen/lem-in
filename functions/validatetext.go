package lemin

import ("strings"
		"log")
func ValidateText(dataArray []string)([]string, []string){
var rooms []string
var links []string
start := false // boolean to check if there is a start room and end room
end := false
for key, value := range dataArray[1:] {
	if strings.HasPrefix(value, "#") {
		if value == "##start" { 
			if !start {
				Colony1.Start = ValidateRoom(dataArray[key+2]) // add Start room to the colony
				Colony1.Start.NumOfAnts = Colony1.Ants // the num of ants to the start room
				Colony1.Start.Occupied = true
				start = true
			} else {
				log.Fatalf("Invalid data format: found duplicate ##start")
			}
		} else if value == "##end" {
			if !end {
				Colony1.End = ValidateRoom(dataArray[key+2]) // add end room to the colony 
				end = true
			} else {
				log.Fatalf("Invalid data format: found duplicate ##end")
			}
		}
	} else if strings.Contains(value, " ") { 
		x := strings.Split(value, " ") // split the data so as to not add the end and start room to the rooms array
		if x[0] == Colony1.End.Name || x[0] == Colony1.Start.Name {
			continue
		} else {
			rooms = append(rooms, value)
		}
	} else if strings.Contains(value, "-") {
		links = append(links, value)
	} else if value == "" {
		continue
	} else {
		log.Fatalf("Invalid data format: weird format")
	}
}
if !start || !end {
	log.Fatalf("Invalid data format: missing start or end")
}
return rooms, links
}