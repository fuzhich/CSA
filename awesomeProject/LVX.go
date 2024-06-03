package main

import (
	"fmt"
)

type Character struct {
	Name       string
	Level      int
	Experience int
	Hp         int
	Atk        int
	Weapen     Weapen
}

type Weapen struct {
	Name string
	Atk  int
}

type Action interface {
	Attack(c1 *Character, c2 *Character)
}

func (c *Character) Attack(c1 *Character, c2 *Character) {
	if c2.Hp-c1.Atk-c1.Weapen.Atk <= 0 {
		c2.Hp = 0
	} else {
		c2.Hp = c2.Hp - c1.Atk - c1.Weapen.Atk
	}
}

type information interface {
	CInformation(c *Character)
}

func (c *Character) CInformation(c1 *Character) {
	fmt.Println(c1.Name)
	fmt.Println(c1.Hp)
}

func main() {
	var wufengjian Weapen
	wufengjian.Name = "无锋剑"
	wufengjian.Atk = 60

	var hand Weapen
	hand.Name = "拳头"
	hand.Atk = 20

	var ying Character
	ying.Name = "荧"
	ying.Level = 1
	ying.Experience = 0
	ying.Hp = 100
	ying.Atk = 20
	ying.Weapen = wufengjian

	var qiuqiuren Character
	qiuqiuren.Name = "丘丘人"
	qiuqiuren.Level = 1
	qiuqiuren.Experience = 0
	qiuqiuren.Hp = 100
	qiuqiuren.Atk = 20
	qiuqiuren.Weapen = hand

	var Aying Action
	Aying = &ying
	Aying.Attack(&ying, &qiuqiuren)

	var Iying information
	Iying = &ying
	Iying.CInformation(&ying)
}
