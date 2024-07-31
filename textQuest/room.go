package main


type Room struct {
	Name        string
	Description string
	Items       map[string]*Item
	Exits       map[string]*Room
}

func (r *Room) getExits() []string {
	exits := []string{}
	for name := range r.Exits {
		exits = append(exits, name)
	}
	return exits
}
