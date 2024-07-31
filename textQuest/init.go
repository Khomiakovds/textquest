package main

func initGame() {
	
	kitchen := &Room{
		Name:        "кухня",
		Description: "ты находишься на кухне, на столе: чай. надо собрать рюкзак и идти в универ. можно пройти - коридор",
		Items:       map[string]*Item{"чай": {"чай", "чай на столе"}},
		Exits:       make(map[string]*Room),
	}
	corridor := &Room{
		Name:        "коридор",
		Description: "ничего интересного. можно пройти - кухня, комната, улица",
		Items:       make(map[string]*Item),
		Exits:       make(map[string]*Room),
	}
	room := &Room{
		Name:        "комната",
		Description: "ты в своей комнате. можно пройти - коридор",
		Items: map[string]*Item{
			"ключи":     {"ключи", "ключи на столе"},
			"конспекты": {"конспекты", "конспекты на столе"},
			"рюкзак":    {"рюкзак", "рюкзак на стуле"},
		},
		Exits: make(map[string]*Room),
	}
	outside := &Room{
		Name:        "улица",
		Description: "на улице весна. можно пройти - домой",
		Items:       make(map[string]*Item),
		Exits:       make(map[string]*Room),
	}


	kitchen.Exits["коридор"] = corridor
	corridor.Exits["кухня"] = kitchen
	corridor.Exits["комната"] = room
	room.Exits["коридор"] = corridor

	
	rooms = map[string]*Room{
		"кухня":   kitchen,
		"коридор": corridor,
		"комната": room,
		"улица":   outside,
	}


	player = Player{
		CurrentRoom: kitchen,
		Inventory:   make(map[string]*Item),
		Backpack:    &Backpack{Items: make(map[string]*Item)},
	}
}
