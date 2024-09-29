package iunit

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

// -------------------------
// Монстр - юнит с автовыбором поведения.
// -------------------------
type Monster struct {
	baseUnit
}

func (m *Monster) SetActions(d Defender) {
	m.blocks = nil
	m.targets = nil

	modes := []string{"A", "D"}

	var mode string

	for i := 0; i < int(m.actions); i++ {
		mode = modes[rand.Intn(len(modes))]
		if mode == "A" {
			m.ChooseAttack(d)
		} else {
			m.ChooseDefence()
		}
	}
}

func (m *Monster) ChooseAttack(d Defender) {
	bodyParts := d.GetBodyParts()
	m.targets = append(m.targets, bodyParts[rand.Intn(len(bodyParts))])
}

func (m *Monster) ChooseDefence() {
	bodyParts := m.GetBodyParts()
	m.targets = append(m.blocks, bodyParts[rand.Intn(len(bodyParts))])
}

// -------------------------
// Создание монстра.
// -------------------------
func NewMonster(name string, icon string, hp, actions int, bps []BodyPart) Monster {
	c := color.New(color.Underline)
	return Monster{
		baseUnit: baseUnit{
			name:      fmt.Sprintf("%v %v", icon, c.Sprintf(name)),
			maxHp:     hp,
			currentHp: hp,
			damage:    1,
			actions:   actions,
			bodyParts: bps,
		},
	}
}
