package Dice

import (
	"math/rand"
)

type D12 struct {
	NumRolled int
	Max       int
	Min       int
}

func (d *D12) Roll() {
	d.NumRolled = rand.Intn(d.Max-d.Min) + d.Min
}

func (d *D12) Init() {
	d.Min = 1
	d.Max = 12
	d.NumRolled = 0
}
