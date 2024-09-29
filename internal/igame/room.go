package igame

import (
	"fmt"
	"rooms/internal/iunit"
	"strconv"

	"github.com/fatih/color"
)

type room struct {
	number  uint
	player  *iunit.Player
	monster iunit.Monster
	round   round
}

func (r *room) start(p *iunit.Player) {
	r.player = p

	for r.monster.IsAlive() && r.player.IsAlive() {
		r.toNextRound()
	}

	if !r.monster.IsAlive() {
		fmt.Println()
		color.Red("%v помер", r.monster.GetName())
	}

	if !r.player.IsAlive() {
		fmt.Println()
		color.Red("%v помер", r.player.GetName())
	}
}

func (r *room) toNextRound() {
	r.round = NewRound(r.round.number + 1)
	fmt.Println(r.round, "\n")

	r.player.SetActions(&r.monster)
	r.monster.SetActions(r.player)

	fmt.Println()

	r.player.Attack(&r.monster)
	if r.monster.IsAlive() {
		r.monster.Attack(r.player)
	}
}

func (r room) String() string {
	return fmt.Sprintf("Комната: %v\nПротивник: %v", r.number, r.monster.String())
}

func NewRoom(number uint) room {
	unitProps := iunit.GetUnit()

	unitHp, _ := strconv.Atoi(unitProps["hp"])
	unitActions, _ := strconv.Atoi(unitProps["actions"])

	monster := iunit.NewMonster(unitProps["name"], unitProps["icon"], unitHp, unitActions, []iunit.BodyPart{iunit.Head, iunit.Chest, iunit.Leg})

	return room{
		number:  number,
		monster: monster,
	}
}
