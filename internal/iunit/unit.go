package iunit

import (
	"fmt"
	"slices"
)

// -------------------------
// Защищающийся.
// -------------------------
type Defender interface {
	IsAlive() bool
	Defend(bp BodyPart, damage int) int
	GetBodyParts() []BodyPart
	GetHp() int
}

// -------------------------
// Базовый параметры юнита. Юнитом являются любые участники битвы.
// -------------------------
type baseUnit struct {
	name             string
	maxHp, currentHp int
	actions, damage  int
	bodyParts        []BodyPart
	blocks           []BodyPart
	targets          []BodyPart
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

func (u *baseUnit) GetBodyPart(bp BodyPart) int {
	return slices.IndexFunc(u.GetBodyParts(), func(item BodyPart) bool {
		return bp == item
	})
}

// Части тела.
func (u *baseUnit) GetBodyParts() []BodyPart {
	return u.bodyParts
}

// Наличие части тела.
func (u *baseUnit) HasBodyPart(bp BodyPart) bool {
	return u.GetBodyPart(bp) != -1
}

// Проверка наличия блока на часте тела.
func (u *baseUnit) underBlock(bp BodyPart) bool {
	i := slices.IndexFunc(u.blocks, func(item BodyPart) bool {
		return bp == item
	})

	return i != -1
}

// Блокирование.
func (u *baseUnit) blockHit(bp BodyPart) {
	i := slices.IndexFunc(u.blocks, func(item BodyPart) bool {
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
func (u *baseUnit) Defend(bp BodyPart, damage int) int {
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
func (u *baseUnit) AttackBodyPart(d Defender, bp BodyPart) int {
	return d.Defend(bp, u.damage)
}

// Полная атака. Зависит от выбранных целей.
func (m *baseUnit) Attack(d Defender) {
	for _, bodyPart := range m.targets {
		if !d.IsAlive() {
			break
		}

		realDamage := m.AttackBodyPart(d, bodyPart)
		if realDamage > 0 {
			fmt.Printf("🟢 %s атакует и попадает по %s\n", m, bodyPart)
		} else {
			fmt.Printf("🔴 %s атакует, но промахивается\n", m)
		}
	}
}
