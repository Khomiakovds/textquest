package main

import (
	"fmt"
	"strings"
	"sort"
)


type Player struct {
	CurrentRoom *Room
	Inventory   map[string]*Item
	Backpack    *Backpack
}

func (p *Player) lookAround() string {
    room := p.CurrentRoom
    items := make([]string, 0, len(room.Items))
    for name := range room.Items {
        items = append(items, name)
    }
    sort.Strings(items) 

    var itemDescription string
    if len(items) > 0 {
        itemDescription = "на столе: " + strings.Join(items, ", ")
    } else {
        itemDescription = "пустая комната"
    }

    exits := strings.Join(room.getExits(), ", ")
    return fmt.Sprintf("%s. %s. можно пройти - %s", room.Description, itemDescription, exits)
}

func (p *Player) goDirection(direction string) string {
    nextRoom, ok := p.CurrentRoom.Exits[direction]
    if !ok {
        return fmt.Sprintf("нет пути в %s", direction)
    }

    if direction == "улица" && !p.isDoorOpen() {
        return "дверь закрыта"
    }

    p.CurrentRoom = nextRoom
    return nextRoom.Description
}

func (p *Player) takeItem(itemName string) string {
    item, ok := p.CurrentRoom.Items[itemName]
    if !ok {
        return "нет такого"
    }

    if _, hasBackpack := p.Backpack.Items["рюкзак"]; !hasBackpack && itemName != "рюкзак" {
        return "некуда класть"
    }

    p.Inventory[itemName] = item
    delete(p.CurrentRoom.Items, itemName)
    return fmt.Sprintf("предмет добавлен в инвентарь: %s", itemName)
}


func (p *Player) wearItem(itemName string) string {
	if itemName != "рюкзак" {
		return fmt.Sprintf("нельзя надеть: %s", itemName)
	}
	if item, ok := p.CurrentRoom.Items[itemName]; ok {
		p.Backpack.Items[itemName] = item  
		delete(p.CurrentRoom.Items, itemName)
		return fmt.Sprintf("вы надели: %s", itemName)
	}
	return "нет такого"
}


func (p *Player) useItem(itemName, target string) string {
    _, ok := p.Inventory[itemName]
    if !ok {
        return fmt.Sprintf("нет предмета в инвентаре - %s", itemName)
    }

    if target == "дверь" && itemName == "ключи" {
        if p.CurrentRoom.Name == "коридор" {
            p.CurrentRoom.Exits["улица"] = rooms["улица"]
            return "дверь открыта"
        }
    }

    return "не к чему применить"
}



func (p *Player) isDoorOpen() bool {
    _, open := rooms["коридор"].Exits["улица"]
    return open
}
