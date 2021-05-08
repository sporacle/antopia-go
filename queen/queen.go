package queen

import "math/rand"
import ants "antopia/ants"

type Queen struct {
	Age int
	Fertility float32
	LifeExpectancy int
	id_count int
	Alive bool
}

func CreateQueen(fertility float32) Queen {
	return Queen{0, fertility, 10950, 0, true}
}

func CreateAnts(queen *Queen) ([]ants.Ant, int, int) {
	if queen.Age > queen.LifeExpectancy {
		queen.Alive = false
		return make([]ants.Ant, 0), 0, 0
	}
	var new_ants []ants.Ant
	var soldierCount int
	var workerCount int
	maxHatch := int(1000.0 * queen.Fertility)
	hatchCount := rand.Intn(maxHatch)
	for hatchCount > 0 {
		hatchCount -= 1
		job := chooseJob()
		if job == "soldier" {
			soldierCount += 1
		} else {
			workerCount += 1
		}
		new_ants = append(new_ants, ants.CreateAnt(queen.id_count, chooseJob()))
		queen.id_count += 1
	}
	queen.Age += 1
	return new_ants, soldierCount, workerCount
}

// ------ HELPERS -------

func chooseJob() string {
	jobRoll := rand.Intn(9)
	switch jobRoll {
	case 1:
		return "soldier"
	default:
		return "worker"
	}	
}