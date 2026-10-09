package builder

import "errors"

var ErrInsufficientPower = errors.New("weapon power overwhelms PowerCore's supply.")
var ErrNoPower = errors.New("no power whatsoever")

type Spaceship interface {
	SetWeapons(weapon string, power int) Weaponry
	SetEngine(engine string)	string
	SetPowerCore(powerLevel int, coreName string) PowerCore
	SetArmor(durability int, armorName string) Armor
}

type ShieldedSpaceship interface {
	SetShield(shield string, cap int) Shield
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