package unit

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/fatih/color"
)

type defender interface {
	isAlive() bool
	defend(bp string, damage int) int
	getBodyParts() []string
	getHp() int
}

// -------------------------
// Базовый параметры юнита. Юнитом являются любые участники битвы.
// -------------------------
type baseUnit struct {
	name             string
	maxHp, currentHp int
	actions, damage  int
	bodyParts        []string
	blocks           []string
	targets          []string
}

func (u *baseUnit) String() string {
	return fmt.Sprintf("%v(%v)", u.name, u.currentHp)
}

func (u *baseUnit) getHp() int {
	return u.currentHp
}

func (u *baseUnit) getBodyPart(name string) int {
	return slices.IndexFunc(u.getBodyParts(), func(item string) bool {
		return item == name
	})
}

// Части тела.
func (u *baseUnit) getBodyParts() []string {
	return u.bodyParts
}

// Наличие части тела.
func (u *baseUnit) hasBodyPart(name string) bool {
	return u.getBodyPart(name) != -1
}

// Проверка наличия блока на часте тела.
func (u *baseUnit) underBlock(name string) bool {
	i := slices.IndexFunc(u.blocks, func(item string) bool {
		return item == name
	})

	return i != -1
}

// Блокирование.
func (u *baseUnit) blockHit(bp string) {
	i := slices.IndexFunc(u.blocks, func(item string) bool {
		return item == bp
	})

	u.blocks = slices.Delete(u.blocks, i, i+1)
}

// Жив?
func (u *baseUnit) isAlive() bool {
	return u.currentHp > 0
}

// Обработка повреждения части тела.
// Возвращается статус успешности атаки.
func (u *baseUnit) defend(bp string, damage int) int {
	if !u.hasBodyPart(bp) {
		return 0
	}

	if u.underBlock(bp) {
		u.blockHit(bp)
		return 0
	}

	return u.getHit(damage)
}

// Подсчет потери очков здоровья.
func (u *baseUnit) getHit(damage int) int {
	u.currentHp -= damage
	if u.currentHp < 0 {
		u.currentHp = 0
	}

	return damage
}

// Атака части тела.
func (u *baseUnit) attackBodyPart(d defender, bodyPart string) int {
	return d.defend(bodyPart, u.damage)
}

// Полная атака. Зависит от выбранных целей.
func (m *baseUnit) attack(d defender) {
	for _, bodyPart := range m.targets {
		if !d.isAlive() {
			break
		}

		realDamage := m.attackBodyPart(d, bodyPart)
		if realDamage > 0 {
			fmt.Println("🟢", m, "атакует и попадает по ", bodyPart)
		} else {
			fmt.Println("🔴", m, "атакует, но промахивается")
		}
	}
}

// -------------------------
// Монстр - юнит с автовыбором поведения.
// -------------------------
type monster struct {
	baseUnit
}

func (m *monster) setActions(d defender) {
	m.blocks = nil
	m.targets = nil

	modes := []string{"A", "D"}

	var mode string

	for i := 0; i < int(m.actions); i++ {
		mode = modes[rand.Intn(len(modes))]
		if mode == "A" {
			m.chooseAttack(d)
		} else {
			m.chooseDefence()
		}
	}
}

func (m *monster) chooseAttack(d defender) {
	bodyParts := d.getBodyParts()
	m.targets = append(m.targets, bodyParts[rand.Intn(len(bodyParts))])
}

func (m *monster) chooseDefence() {
	bodyParts := m.getBodyParts()
	m.targets = append(m.blocks, bodyParts[rand.Intn(len(bodyParts))])
}

// -------------------------
// Создание монстра.
// -------------------------
func NewMonster(name string, icon string, hp, actions int) monster {
	c := color.New(color.Underline)
	return monster{
		baseUnit: baseUnit{
			name:      fmt.Sprintf("%v %v", icon, c.Sprintf(name)),
			maxHp:     hp,
			currentHp: hp,
			damage:    1,
			actions:   actions,
			bodyParts: []string{"Голова", "Торс", "Ноги"},
		},
	}
}
