package iunit

import (
	"fmt"
	"os"
	"rooms/internal/ihelp"

	"github.com/fatih/color"
)

// -------------------------
// Игрок - юнит под управлением игрока.
// -------------------------
type Player struct {
	baseUnit
}

// Изменить кол-во очков жизней. Может быть, как в плюс, так и в минус.
func (u *Player) UpdateHp(value int) {
	u.currentHp += value
	if u.currentHp < 0 {
		u.currentHp = 0
	}

	if u.currentHp > u.maxHp {
		u.currentHp = u.maxHp
	}
}

// Выбор действий.
func (u *Player) SetActions(d Defender) {
	u.blocks = nil
	u.targets = nil

	var mode string

	for i := 0; i < int(u.actions); i++ {
		mode = ihelp.Radio("Что делаем?", []string{"Атака", "Защита"})

		switch mode {
		case "Атака":
			u.ChooseAttack(d)
		case "Защита":
			u.ChooseDefence()
		default:
			os.Exit(3)
		}
	}
}

func (u *Player) ChooseAttack(d Defender) {
	answer := ihelp.Radio("Куда атакуем?", d.GetBodyParts())
	u.targets = append(u.targets, answer)
}

func (u *Player) ChooseDefence() {
	answer := ihelp.Radio("Что защищаем?", u.GetBodyParts())
	u.blocks = append(u.blocks, answer)
}

// -------------------------
// Создание игрока.
// -------------------------
func NewPlayer(hp, actions int) Player {
	c := color.New(color.Underline)
	return Player{
		baseUnit: baseUnit{
			name:      fmt.Sprintf("💩 %v", c.Sprintf("Игрок")),
			maxHp:     hp,
			currentHp: hp,
			damage:    2,
			actions:   actions,
			bodyParts: []string{"Голова", "Торс", "Ноги"},
		},
	}
}
