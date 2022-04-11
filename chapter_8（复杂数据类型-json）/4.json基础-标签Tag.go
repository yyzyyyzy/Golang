package main

import (
	"encoding/json"
	"fmt"
)

type Skill struct {
	SkillName string `json:"Penis"`
	Level     int    `json:"Pussy"`
}

type Actor struct {
	Name   string
	Age    int
	Skills []Skill
}

func main() {
	actor := Actor{Name: "LZK", Age: 18, Skills: []Skill{
		{SkillName: "Java", Level: 0},
		{SkillName: "Python", Level: 8},
		{SkillName: "Golang", Level: 5},
	}}

	result, _ := json.Marshal(actor)
	fmt.Printf("%s", result)
}
