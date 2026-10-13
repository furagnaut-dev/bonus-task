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

func (d *Director) Warship() (*builder.Ship, error) {
	return d.builder.
		SetPowerCore(19, "Malefic").
		SetEngine("Ion Drive").
		SetArmor(500, "Titanium").
		SetWeapons("Plasma Cannon", 8).
		SetWeapons("Railgun", 10).
		Build()
}

func ShieldedWarship(b builder.ShieldedSpaceship) (*builder.Ship, error) {
	return b.SetShield("Aegis", 300).
		SetPowerCore(19, "Malefic").
		SetEngine("Ion Drive").
		SetArmor(500, "Titanium").
		SetWeapons("Plasma Cannon", 8).
		Build()
}
