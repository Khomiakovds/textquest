package main

import (
	//"fmt"
	"strings"
)

func handleCommand(command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "Пожалуйста, введите команду."
	}

	cmd := parts[0]
	param1 := ""
	param2 := ""
	if len(parts) > 1 {
		param1 = parts[1]
	}
	if len(parts) > 2 {
		param2 = parts[2]
	}

	switch cmd {
	case "осмотреться":
		return player.lookAround()
	case "идти":
		return player.goDirection(param1)
	case "взять":
		return player.takeItem(param1)
	case "надеть":
		return player.wearItem(param1)
	case "применить":
		return player.useItem(param1, param2)
	default:
		return "неизвестная команда"
	}
}
