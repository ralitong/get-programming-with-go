package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor{
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())

	var k kelvin = 294.0

	otherSensor := func() kelvin {
		return k
	}

	fmt.Println(otherSensor())
	k++
	fmt.Println(otherSensor())

	newCalibrate := 10
	sensor = calibrate(realSensor, kelvin(newCalibrate))

	fmt.Println(sensor())
	newCalibrate += 10
	fmt.Println(sensor())

	sensor = calibrate(fakeSensor, kelvin(newCalibrate))
	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
	
}