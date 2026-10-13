package builder

import "errors"

var ErrInsufficientPower = errors.New("weapon power overwhelms PowerCore's supply.")
var ErrNoPower = errors.New("no power whatsoever")

type Spaceship interface {
	SetWeapons(weapon string, power int) Spaceship
	SetEngine(engine string)	Spaceship
	SetPowerCore(powerLevel int, coreName string) Spaceship
	SetArmor(durability int, armorName string) Spaceship
	Build() (*Ship, error)
}

type ShieldedSpaceship interface {
	Spaceship
	SetShield(shield string, cap int) Spaceship
}

type Shield struct {
	Name string
	Capacity int
}

type PowerCore struct {
	Power int
	Core string
}

type Weaponry struct {
	Name string
	Power int
}

type Armor struct {
	Durability int
	Name string
}

type Ship struct {
	Weapons   []Weaponry
	Engine    string
	PowerCore PowerCore
	Armor     Armor
	Shield    *Shield
}

func ValidateWeaponPower(weapon []Weaponry, core PowerCore) error {
	powerLevel := 0
	for _, weaponry := range weapon {
		powerLevel += weaponry.Power
		if weaponry.Power <= 0 {
			return ErrNoPower
		}
	}
	if core.Power < 0 {
		return ErrNoPower
	}
	if powerLevel > core.Power {
		return ErrInsufficientPower
	}

	return nil
}