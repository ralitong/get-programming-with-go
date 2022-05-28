package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal interface {
	eat() string
	sleep() string
	move() string
}

type mammal struct {
	name string
	food string
}

func (m mammal) String() string{
	return m.name
}

func (m mammal) eat() string{
	return fmt.Sprintf("%v is eating %v", m, m.food)
}

func (m mammal) sleep() string{
	return fmt.Sprintf("%v is sleeping", m)
}

func (m mammal) move() string{
	return fmt.Sprintf("%v is moving", m)
}

type sol int

func (s sol) hours() int {
	return int(s) * 3
}

func pickRandomAnimalAction(a animal) {
	random := rand.Intn(1)

	if(random == 0) {
		fmt.Println(a.eat()) 
	} else {
		fmt.Println(a.move())
	}
}

func letTheAnimalSleep(a animal) {
	fmt.Println(a.sleep())
}

type clock struct {
	hour int
	night bool
}

func tick(c clock) clock{
	return clock{c.hour+1, c.night}
}

func cycle(c clock) clock {
	if c.hour % 12 == 0 {
		return clock{c.hour, !c.night}
	}
	return c
}

func main() {
	cycleClock := clock{}
	
	sols := sol(24)
	animals := []animal{
		mammal{"dog", "royal canin"},
		mammal{"cow", "grass"},
		mammal{"cat", "whiskas"},
	}

	for i := 0; i < sols.hours(); i++ {
		randomAnimal := rand.Intn(len(animals))
		if(!cycleClock.night) {
			pickRandomAnimalAction(animals[randomAnimal])
		} else {
			letTheAnimalSleep(animals[randomAnimal])
		}
		cycleClock = tick(cycleClock)
		cycleClock = cycle(cycleClock)
		time.Sleep(time.Second)
	}


}