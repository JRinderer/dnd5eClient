package Dice

import (
	"math/rand"
)

type D10 struct {
	NumRolled int
	Max       int
	Min       int
}

func (d *D10) Roll() {
	d.NumRolled = rand.Intn(d.Max-d.Min) + d.Min
}

func (d *D10) Init() {
	d.Min = 1
	d.Max = 10
	d.NumRolled = 0
}
