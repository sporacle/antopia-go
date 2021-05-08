package colony

import "fmt"
import "math"
import "math/rand"
import hab "antopia/habitat"
import qn "antopia/queen"
import ants "antopia/ants"

type Colony struct {
	Name string

	habitat *hab.Habitat
	queen *qn.Queen
	ColonyAnts map[int]ants.Ant

	Age int
	// Food level in milligrams
	FoodLevel int
	dayNum int
}

func CreateColony(name string, habitat *hab.Habitat, queen *qn.Queen) Colony {
	fmt.Printf("New colony called %v with food abundance of %.1f and danger level of %.1f. The queen has a fertility of %.1f.\n",
	 name, habitat.FoodLevel, habitat.DangerLevel, queen.Fertility)
	return Colony{name, habitat, queen, make(map[int]ants.Ant), 0, 0, 0}
}

func PassDay(colony *Colony) {
	colony.dayNum += 1
	fmt.Printf("\n")
	fmt.Printf("------------------------------------\n")
	fmt.Printf("--- Day %d goes by in the colony ---\n", colony.dayNum)
	fmt.Printf("------------------------------------\n")

	foodLevel, dangerLevel := hab.GetHabitatStats(colony.habitat)
		fmt.Printf("food: %.2f danger: %.2f\n", foodLevel, dangerLevel)
	// Hatch some new ants
	new_ants, soldierCount, workerCount := qn.CreateAnts(colony.queen)
	combineAnts(colony, new_ants)
	fmt.Printf("The queen hatched %d new workers and %d new soldiers totalling %d ants.\n", workerCount, soldierCount, len(new_ants))
	fmt.Printf("The colony has a total population of %d\n", len(colony.ColonyAnts))

	// Have the ants do their jobs in the colony
	var foodGathered int
	var soldiersDied int
	var workersDied int

	for k, v := range colony.ColonyAnts {
		surivived, food := ants.PerformJob(&v, foodLevel, dangerLevel)
		if (surivived) {
			foodGathered += food
		} else {
			switch v.Job {
			case "worker":
				workersDied += 1
			case "soldier":
				soldiersDied += 1
			}
			delete(colony.ColonyAnts, k)
		}
	}
	fmt.Printf("After a day of duties %dmg of food was gathered and %d workers and %d soldiers died.\n", foodGathered, workersDied, soldiersDied)
	colony.FoodLevel += foodGathered
	feedColony(colony)
}

// ------ HELPERS -------

func combineAnts(colony *Colony, new_ants []ants.Ant) {
	for _, v := range new_ants {
		colony.ColonyAnts[v.Id] = v
	}
}

func feedColony(colony *Colony) {
	initialFood := colony.FoodLevel
	// The queen eats 300mg first and then the colony ants will eat 2mg each.
	if colony.queen.Alive {
		colony.FoodLevel -= 300
	}
	population := len(colony.ColonyAnts)
	colony.FoodLevel -= population * 2
	if colony.FoodLevel < 0 {
		// Some ants will die of starvation depending on the food deficit
		hungryAnts := int(math.Abs(float64(colony.FoodLevel / 2)))
		starveColony(colony, hungryAnts)
		colony.FoodLevel = 0
	}
	fmt.Printf("%dmg of food was eaten by the colony, %dmg remains.\n", initialFood - colony.FoodLevel, colony.FoodLevel)
}

func starveColony(colony *Colony, hungryAnts int) {
	// 10% of the ants who didn't eat will die
	toDie := hungryAnts / 10
	var soldiersDied int
	var workersDied int
	fmt.Printf("The colony is starving\n")
	for i := 1; i < toDie; i++ {
		ids := getIds(colony)
		starving_ant := ids[rand.Intn(len(ids))]
		switch colony.ColonyAnts[starving_ant].Job {
			case "worker":
				workersDied += 1
			case "soldier":
				soldiersDied += 1
			}
		delete(colony.ColonyAnts, starving_ant)    	
	}
	fmt.Printf("%d workers and %d soldiers starved to death.\n", workersDied, soldiersDied)
}

func getIds(colony *Colony) []int {
	keys := make([]int, len(colony.ColonyAnts))
	i := 0
	for k := range colony.ColonyAnts {
	    keys[i] = k
	    i++
	}
	return keys
}