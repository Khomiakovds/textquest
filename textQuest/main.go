package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var player Player
var rooms map[string]*Room

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в игру! Введите команду:")
	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "выход" {
			break
		}
		fmt.Println(handleCommand(command))
	}
	fmt.Println("Игра завершена. Спасибо за игру!")
	
}
