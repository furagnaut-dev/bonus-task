package concretebuilders

import (
	"errors"
	"spaceship/builder"
)

var ErrNoNecessaryComponent = errors.New("has no necessary component")

type SpaceshipBuilder struct {
	weapons []builder.Weaponry
	engine string
	powerCore builder.PowerCore
	armor builder.Armor
	shield builder.Shield
	hasEngine bool
	hasCore bool
	hasArmor bool
}

func NewSpaceship() *SpaceshipBuilder {
	return &SpaceshipBuilder{}
}

func (s *SpaceshipBuilder) SetWeapons(name string, power int) builder.Weaponry {
	weapon := builder.Weaponry{Name: name, Power: power}
	s.weapons = append(s.weapons, weapon)
	return weapon
}

func (s *SpaceshipBuilder) SetEngine(engine string) string {
	s.engine = engine
	s.hasEngine = true
	return engine
}

func (s *SpaceshipBuilder) SetPowerCore(powerLevel int, coreName string) builder.PowerCore {
	core := builder.PowerCore{Power: powerLevel, Core: coreName}
	s.powerCore = core
	s.hasCore = true
	return core
}

func (s *SpaceshipBuilder) SetArmor(durability int, armorName string) builder.Armor {
	armor := builder.Armor{Durability: durability, Name: armorName}
	s.armor = armor
	s.hasArmor = true
	return armor
}

func (s *SpaceshipBuilder) SetShield(shield string, cap int) builder.Shield {
	shielder := builder.Shield{Name: shield, Capacity: cap}
	s.shield = shielder
	return shielder
}

func (s *SpaceshipBuilder) Build() (*SpaceshipBuilder, error) {
	if !s.hasArmor || !s.hasCore || !s.hasEngine {
		return nil, ErrNoNecessaryComponent
	}
	if err := builder.ValidateWeaponPower(s.weapons, s.powerCore); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SpaceshipBuilder) Engine() string {
	return s.engine
}

func (s *SpaceshipBuilder) PowerCore() builder.PowerCore {
	return s.powerCore
}

func (s *SpaceshipBuilder) Armor() builder.Armor {
	return s.armor
}

func (s *SpaceshipBuilder) Weapons() []builder.Weaponry {
	return s.weapons
}

func (s *SpaceshipBuilder) Shield() builder.Shield {
	return s.shield
}
