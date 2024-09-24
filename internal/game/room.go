package game

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/fatih/color"
)

type room struct {
	number  uint
	player  *player
	monster monster
	round   round
}

func (r *room) start(p *player) {
	r.player = p

	for r.monster.isAlive() && r.player.isAlive() {
		r.toNextRound()
	}

	if !r.monster.isAlive() {
		fmt.Println()
		color.Red("%v помер", r.monster.name)
	}

	if !r.player.isAlive() {
		fmt.Println()
		color.Red("%v помер", r.player.name)
	}
}

func (r *room) toNextRound() {
	r.round = NewRound(r.round.number + 1)
	fmt.Println(r.round, "\n")

	r.player.setActions(&r.monster)
	r.monster.setActions(r.player)

	fmt.Println()

	r.player.attack(&r.monster)
	if r.monster.isAlive() {
		r.monster.attack(r.player)
	}
}

func (r room) String() string {
	return fmt.Sprintf("Комната: %v\nПротивник: %v", r.number, r.monster.String())
}

func NewRoom(number uint) room {
	unitProps := units[rand.Intn(len(units))]

	unitHp, _ := strconv.Atoi(unitProps["hp"])
	unitActions, _ := strconv.Atoi(unitProps["actions"])

	return room{
		number:  number,
		monster: NewMonster(unitProps["name"], unitProps["icon"], unitHp, unitActions),
	}
}
