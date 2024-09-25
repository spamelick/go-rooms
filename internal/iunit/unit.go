package iunit

import (
	"fmt"
	"slices"
)

type Defender interface {
	IsAlive() bool
	Defend(bp string, damage int) int
	GetBodyParts() []string
	GetHp() int
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

func (u *baseUnit) GetName() string {
	return u.name
}

func (u *baseUnit) GetHp() int {
	return u.currentHp
}

func (u *baseUnit) GetBodyPart(name string) int {
	return slices.IndexFunc(u.GetBodyParts(), func(item string) bool {
		return item == name
	})
}

// Части тела.
func (u *baseUnit) GetBodyParts() []string {
	return u.bodyParts
}

// Наличие части тела.
func (u *baseUnit) HasBodyPart(name string) bool {
	return u.GetBodyPart(name) != -1
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
func (u *baseUnit) IsAlive() bool {
	return u.currentHp > 0
}

// Обработка повреждения части тела.
// Возвращается статус успешности атаки.
func (u *baseUnit) Defend(bp string, damage int) int {
	if !u.HasBodyPart(bp) {
		return 0
	}

	if u.underBlock(bp) {
		u.blockHit(bp)
		return 0
	}

	return u.GetHit(damage)
}

// Подсчет потери очков здоровья.
func (u *baseUnit) GetHit(damage int) int {
	u.currentHp -= damage
	if u.currentHp < 0 {
		u.currentHp = 0
	}

	return damage
}

// Атака части тела.
func (u *baseUnit) AttackBodyPart(d Defender, bodyPart string) int {
	return d.Defend(bodyPart, u.damage)
}

// Полная атака. Зависит от выбранных целей.
func (m *baseUnit) Attack(d Defender) {
	for _, bodyPart := range m.targets {
		if !d.IsAlive() {
			break
		}

		realDamage := m.AttackBodyPart(d, bodyPart)
		if realDamage > 0 {
			fmt.Println("🟢", m, "атакует и попадает по ", bodyPart)
		} else {
			fmt.Println("🔴", m, "атакует, но промахивается")
		}
	}
}
