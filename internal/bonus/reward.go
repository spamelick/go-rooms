package bonus

type talent interface {
	getName() string
	apply(p *player)
	getDescription() string
}

// -------------------------
// Базовая награда.
// -------------------------
type baseTalent struct {
	name        string
	description string
	handler     func(*player, map[string]string)
}

func (r *baseTalent) getName() string {
	return r.name
}

func (r *baseTalent) getDescription() string {
	return r.description
}

// -------------------------
// Автолечение после прохождения комнаты.
// -------------------------
type healAfterWinTalent struct {
	baseTalent
}

func (t *healAfterWinTalent) apply(p *player) {

}

func NewTalent(name, description string, handler func(*player)) baseTalent {
	return baseTalent{
		name:        name,
		description: description,
		handler:     handler,
	}
}

// -------------------------
// Бонусы.
// -------------------------
func healSkill(p *player, params map[string]string) {
	p.updateHp(int(params["value"]))
}

// бонус к урону
// лечение
// шанс двойного атаки
// шанс двойного урона
// шанс двойного урона в голову
// автолечение 1хп
// рост атаки c каждым раунд
