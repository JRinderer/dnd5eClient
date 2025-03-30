package main

import (
	"DND5eClient/Dice"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	d := Dice.D20{}
	d.Init()
	d.Roll()
	fmt.Println(d.NumRolled)
	d12 := Dice.D12{}
	d12.Init()
	d12.Roll()
	fmt.Println(d12.NumRolled)
	d10 := Dice.D10{}
	d10.Init()
	d10.Roll()
	fmt.Println(d10.NumRolled)

	d8 := Dice.D8{}
	d8.Init()
	d8.Roll()
	fmt.Println(d8.NumRolled)

	d6 := Dice.D6{}
	d6.Init()
	d6.Roll()
	fmt.Println(d6.NumRolled)

	d4 := Dice.D4{}
	d4.Init()
	d4.Roll()
	fmt.Println(d4.NumRolled)
}
