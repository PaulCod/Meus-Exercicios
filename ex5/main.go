package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

func main() {
	speed := 5
	batteryDrain := 2
	car := newCar(speed, batteryDrain)
	car = drive(car)

	distance := 100
	raceTrack := newTrack(distance)

	finish := carFinish(car, raceTrack)
	fmt.Println(finish)
}

func newCar(speed, batteryDrain int) Car {
	batteryChange := 100
	return Car{speed: speed, battery: batteryChange, batteryDrain: batteryDrain, distance: 0}
}

func newTrack(distance int) Track {
	return Track{distance: distance}
}

func drive(car Car) Car {
	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

func carFinish(car Car, raceTrack Track) bool {
	for car.battery != 0 {
		car = drive(car)
	}

	return car.distance >= raceTrack.distance
}
