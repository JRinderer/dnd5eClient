package Dice

import (
	"math/rand"
)

type D8 struct {
	NumRolled int
	Max       int
	Min       int
}

func (d *D8) Roll() {
	d.NumRolled = rand.Intn(d.Max-d.Min) + d.Min
}

func (d *D8) Init() {
	d.Min = 1
	d.Max = 8
	d.NumRolled = 0
}
