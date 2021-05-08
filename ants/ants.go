package ants

import "math"
import "math/rand"

type Ant struct {
	Id int
	age int
	life_expectancy int
	Job string
}

func CreateAnt(id int, job string) Ant {
	ant := Ant{id, 0, 0, job}
	switch job {
	case "worker":
		ant.life_expectancy = 1000
	case "soldier":
		ant.life_expectancy = 35
	}
	return ant
}

func PerformJob(ant *Ant, foodLevel float32, dangerLevel float32) (bool, int) {
	ant.age += 1
	if ant.age > ant.life_expectancy {
		return false, 0
	}
	switch ant.Job {
	case "worker":
		return performWorkerJob(ant, foodLevel, dangerLevel)
	case "soldier":
		return performSoldierJob(ant, foodLevel, dangerLevel), 0
	}
	return false, 0
}

// ------ HELPERS -------

func performWorkerJob(ant *Ant, foodLevel float32, dangerLevel float32) (bool, int) {

	// Will the ant survive the day?
	dangerRoll := rand.Intn(1000)
	workerDanger := 10
	realDanger := int(float32(workerDanger) * dangerLevel)
	if dangerRoll <= realDanger {
		return false, 0
	}

	var foodGathered int;
	// Will food be found?
	foundRoll := rand.Intn(1000)
	// Base percent chance of finding food is 5%
	if foundRoll < int(math.Round(float64(float32(50.0 * foodLevel)))) {
		// An ant can carry 50-150 mg of food back to the colony
		foodGathered = 50 + rand.Intn(100)
	}
	// Time to head back
	return true, foodGathered

}

func performSoldierJob(ant *Ant, foodLevel float32, dangerLevel float32) bool {
	// Will the ant survive the day?
	dangerRoll := rand.Intn(1000)
	soldierDanger := 100
	realDanger := int(float32(soldierDanger) * dangerLevel)
	if dangerRoll <= realDanger {
		return false
	}
	return true
}