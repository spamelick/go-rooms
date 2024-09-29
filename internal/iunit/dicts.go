package iunit

import (
	"fmt"
	"math/rand"
)

// -------------------------
// Справочник монстров.
// -------------------------
var units []map[string]string = []map[string]string{
	{
		"icon":    "🐗",
		"name":    "Кабан",
		"hp":      "3",
		"actions": "1",
	},
	{
		"icon":    "🐸",
		"name":    "Липкий жаб",
		"hp":      "1",
		"actions": "3",
	},
	{
		"icon":    "🐢",
		"name":    "Панцирь",
		"hp":      "5",
		"actions": "1",
	},
}

func GetUnit() map[string]string {
	return units[rand.Intn(len(units))]
}

// -------------------------
// Части тела.
// -------------------------
type BodyPart int

const (
	Unknown BodyPart = iota
	Head
	Chest
	Leg
	Tail
)

var bodyParts map[BodyPart]string = map[BodyPart]string{
	Unknown: "-",
	Head:    "Голова",
	Chest:   "Грудь",
	Leg:     "Ноги",
	Tail:    "Хвост",
}

func (bp BodyPart) String() string {
	if val, ok := bodyParts[bp]; ok {
		return val
	}

	return fmt.Sprintf("BodyPart(%q)", int(bp))
}
