package main


func main() {
	type celsius float64
	const degress = 20
	var temperature celsius = degress
	temperature += 10
	
	var warmUp float64 = 10
	temperature += celsius(warmUp)
}