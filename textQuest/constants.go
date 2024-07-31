package main



var Commands = map[string]string{
	"exit":        "выход",
	"look_around": "осмотреться",
	"go":          "идти",
	"take":        "взять",
	"wear":        "надеть",
	"use":         "применить",
}

var Messages = map[string]string{
	"welcome":              "Добро пожаловать в игру! Введите команду:",
	"goodbye":              "Игра завершена. Спасибо за игру!",
	"enter_command":        "Пожалуйста, введите команду.",
	"unknown_command":      "неизвестная команда",
	"door_closed":          "дверь закрыта",
	"no_way":               "нет пути в %s",
	"nowhere_to_put":       "некуда класть",
	"no_such_item":         "такого предмета нет",
	"item_added":           "предмет %s добавлен в рюкзак",
	"item_worn":            "вы надели %s",
	"cannot_wear":          "нельзя надеть %s",
	"no_item_in_inventory": "у вас нет предмета %s",
	"door_opened":          "дверь открыта",
	"nothing_to_use_on":    "не на что применить",
	"kitchen_have_items":   "ты находишься на кухне, на столе: чай. надо собрать рюкзак и идти в универ. можно пройти - коридор",
	"kitchen_no_items":     "ты находишься на кухне, на столе: чай. надо собрать рюкзак и идти в универ. можно пройти - коридор",
}
