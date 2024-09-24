package unit

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// -------------------------
// Игрок - юнит под управлением игрока.
// -------------------------
type player struct {
	baseUnit
}

// Изменить кол-во очков жизней. Может быть, как в плюс, так и в минус.
func (u *player) updateHp(value int) {
	u.currentHp += value
	if u.currentHp < 0 {
		u.currentHp = 0
	}

	if u.currentHp > u.maxHp {
		u.currentHp = u.maxHp
	}
}

// Выбор действий.
func (u *player) setActions(d defender) {
	u.blocks = nil
	u.targets = nil

	var mode string

	for i := 0; i < int(u.actions); i++ {
		mode = radio("Что делаем?", []string{"Атака", "Защита"})

		switch mode {
		case "Атака":
			u.chooseAttack(d)
		case "Защита":
			u.chooseDefence()
		default:
			os.Exit(3)
		}
	}
}

func (u *player) chooseAttack(d defender) {
	answer := radio("Куда атакуем?", d.getBodyParts())
	u.targets = append(u.targets, answer)
}

func (u *player) chooseDefence() {
	answer := radio("Что защищаем?", u.getBodyParts())
	u.blocks = append(u.blocks, answer)
}

// -------------------------
// Создание игрока.
// -------------------------
func NewPlayer(hp, actions int) player {
	c := color.New(color.Underline)
	return player{
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
