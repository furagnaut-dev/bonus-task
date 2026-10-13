package main

import (
	"fmt"
	concretebuilders "spaceship/concreteBuilders"
	"spaceship/director"
)


func main() {
	ship, err := director.NewDirector(concretebuilders.NewSpaceship()).Warship()
	fmt.Printf("Director warship: %+v err=%v\n", ship, err)

	ship, err = director.ShieldedWarship(concretebuilders.NewSpaceship())
	fmt.Printf("Shielded warship: shield=%+v err=%v\n", *ship.Shield, err)

	_, err = concretebuilders.NewSpaceship().
		SetPowerCore(5, "Tiny").SetEngine("Ion").SetArmor(10, "Tin").
		SetWeapons("Death Ray", 50).Build()
	fmt.Println("Overpowered ship:", err)
}