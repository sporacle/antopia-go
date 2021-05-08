package main

import "time"
import "math/rand"
import hab "antopia/habitat"
import col "antopia/colony"
import qn "antopia/queen"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var name string
	var dangerLevel float32
	var foodLevel float32
	var queenFertility float32

	// Set the configuration variables of the Colony
	name = "First Colony"
	dangerLevel = 0.5
	foodLevel = 0.8
	queenFertility = 0.8

	habitat := hab.CreateHabitat(foodLevel, dangerLevel)
	queen := qn.CreateQueen(queenFertility)
	colony := col.CreateColony(name, &habitat, &queen)

	count := 0
	col.PassDay(&colony)
	for len(colony.ColonyAnts) > 0 {
		count += 1
		// time.Sleep(100 * time.Millisecond)
		col.PassDay(&colony)
	}
}