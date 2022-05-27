package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct {}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(1))
}

type starship struct {
	laser
}

type rover string

func (r rover) talk() string {
	return strings.Repeat(string(r), int(2))
}

func main() {
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

	r := rover("whir")
	shout(r)
}