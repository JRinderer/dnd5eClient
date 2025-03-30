package Dice

import (
	"math/rand"
)

type D4 struct {
	NumRolled int
	Max       int
	Min       int
}

func (d *D4) Roll() {
	d.NumRolled = rand.Intn(d.Max-d.Min) + d.Min
}

func (d *D4) Init() {
	d.Min = 1
	d.Max = 4
	d.NumRolled = 0
}
