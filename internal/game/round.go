package game

import "fmt"

type round struct {
	number uint
}

func NewRound(number uint) round {
	return round{
		number: number,
	}
}

func (r round) String() string {
	return fmt.Sprintf("\n🔷 Раунд #%v", r.number)
}
