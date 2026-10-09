package director

import "spaceship/builder"

type Director struct {
	builder builder.Spaceship
}

func NewDirector(b builder.Spaceship) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) GetSpaceship() {
	d.builder.SetPowerCore(19, "Malefic")
}