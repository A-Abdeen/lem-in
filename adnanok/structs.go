package lemin

var (
	Colony1      Colony
	ColonyMatrix [][]Room
	XMax         int
	YMax         int
	TotalPaths   string
)

type Colony struct {
	Ants       int
	Rooms      []Room
	Start      Room
	End        Room
	Paths      [][]string
	AntTracker []Ant
	Matrix     [][]Room
}

type Room struct {
	Name      string
	X         int
	Y         int
	Links     []string
	NewLinks  []FinalLinks
	Occupied  bool
	Occupier  int
	NumOfAnts int
}
type FinalLinks struct {
	ForwardLinks  string
	RoomsToTheEnd int
}

type Ant struct {
	Number   int
	Location Room
}
