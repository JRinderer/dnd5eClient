package Dice

import (
	"math/rand"
)

type D20 struct {
	NumRolled int
	Max       int
	Min       int
}

func (d *D20) Roll() {
	d.NumRolled = rand.Intn(d.Max-d.Min) + d.Min
}

func (d *D20) Init() {
	d.Min = 1
	d.Max = 20
	d.NumRolled = 0
}
