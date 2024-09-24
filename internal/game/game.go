package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/fatih/color"
)

// Игра.
type game struct {
	seed   int
	room   room
	player player
}

func startGame() {
	game := NewGame()

	color.Cyan("Вы входите в 🏰. Обратной дороги нет...")

	game.start()

	color.Red("💀 Вот и сказочке конец.")
	fmt.Println()
}

func confirmStartGame() {
	reader := bufio.NewReader(os.Stdin)
	input("Начинаем?", reader)
	fmt.Println()
	startGame()
}

// Инициализация игры.
// Если передан seed игры, то... создаем с сидом.
func NewGame(seed ...int) game {
	if len(seed) == 1 {
		return loadGame(seed[0])
	}

	player := NewPlayer(5, 2)

	g := game{
		seed:   rand.Int(),
		player: player,
	}

	return g
}

// Начать игру.
func (g *game) start() {
	for g.player.isAlive() {
		fmt.Println()
		g.startNextRoom()
	}
}

// Активация след. комнаты.
// Либо первой, если нет текущей.
func (g *game) startNextRoom() {
	g.room = NewRoom(g.getNextRoomNumber())
	fmt.Println("-------------------")
	fmt.Println(g.room)
	fmt.Println("-------------------")
	g.room.start(&g.player)
}

func (g game) getNextRoomNumber() uint {
	return g.room.number + 1
}

// Загрузка игры из сохранения.
// Пока только в фантазиях.
func loadGame(_ int) game {
	return NewGame()
}
