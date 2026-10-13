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
	shield *builder.Shield
	hasEngine bool
	hasCore bool
	hasArmor bool
}

func NewSpaceship() *SpaceshipBuilder { return &SpaceshipBuilder{} }

func (s *SpaceshipBuilder) SetWeapons(name string, power int) builder.Spaceship {
	s.weapons = append(s.weapons, builder.Weaponry{Name: name, Power: power})
	return s
}

func (s *SpaceshipBuilder) SetEngine(engine string) builder.Spaceship {
	s.engine, s.hasEngine = engine, true
	return s
}

func (s *SpaceshipBuilder) SetPowerCore(powerLevel int, coreName string) builder.Spaceship {
	s.powerCore, s.hasCore = builder.PowerCore{Power: powerLevel, Core: coreName}, true
	return s
}

func (s *SpaceshipBuilder) SetArmor(durability int, armorName string) builder.Spaceship {
	s.armor, s.hasArmor = builder.Armor{Durability: durability, Name: armorName}, true
	return s
}

func (s *SpaceshipBuilder) SetShield(name string, capacity int) builder.Spaceship {
	s.shield = &builder.Shield{Name: name, Capacity: capacity}
	return s
}

func (s *SpaceshipBuilder) Build() (*builder.Ship, error) {
	if !s.hasArmor || !s.hasCore || !s.hasEngine {
		return nil, ErrNoNecessaryComponent
	}
	if err := builder.ValidateWeaponPower(s.weapons, s.powerCore); err != nil {
		return nil, err
	}
	return &builder.Ship{
		Weapons: s.weapons, Engine: s.engine, PowerCore: s.powerCore,
		Armor: s.armor, Shield: s.shield,
	}, nil
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

func (s *SpaceshipBuilder) Shield() *builder.Shield {
	return s.shield
}
