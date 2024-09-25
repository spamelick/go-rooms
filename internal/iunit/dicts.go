package iunit

import "math/rand"

// Справочник монстров.
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
