package Dice

import (
	"math/rand"
)

type D6 struct {
	NumRolled int
	Max       int
	Min       int
}

func (d *D6) Roll() {
	d.NumRolled = rand.Intn(d.Max-d.Min) + d.Min
}

func (d *D6) Init() {
	d.Min = 1
	d.Max = 6
	d.NumRolled = 0
}
