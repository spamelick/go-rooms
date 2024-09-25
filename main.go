package main

import (
	"bufio"
	"fmt"
	"os"
	"rooms/internal/igame"
	"rooms/internal/ihelp"
)

func confirm() bool {
	reader := bufio.NewReader(os.Stdin)
	ihelp.Input("Начинаем?", reader)
	fmt.Println()

	return true
}

func main() {

	fmt.Println("\nДобро пожаловать в игру!")
	fmt.Println("Цель игры: пройти как можно больше комнат.\n")

	if confirm() {
		igame.StartGame()
	}
}
