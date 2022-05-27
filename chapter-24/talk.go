package main

import (
	"fmt"
	"strings"
)

type martian struct {}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(1))
}

var t interface {
	talk() string
}


func main() {

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

}