package habitat

type Habitat struct {
	FoodLevel float32
	DangerLevel float32

	season_day_count int
}

func CreateHabitat(FoodLevel float32, DangerLevel float32) Habitat {
	return Habitat{FoodLevel, DangerLevel, 0}
}

func GetHabitatStats(habitat *Habitat) (float32, float32) {
	if habitat.season_day_count == 360 {
		habitat.season_day_count = 0
	}
	habitat.season_day_count += 1
	if habitat.season_day_count >= 0 && habitat.season_day_count < 90 {  // Winter
		return habitat.FoodLevel * 0.2, habitat.DangerLevel * 1.5
	} else if habitat.season_day_count >= 90 && habitat.season_day_count < 180 {  // Spring
		return habitat.FoodLevel * 1.2, habitat.DangerLevel * 1.0
	} else if habitat.season_day_count >= 180 && habitat.season_day_count < 270 {  // Summer
		return habitat.FoodLevel * 1.0, habitat.DangerLevel * 1.0
	} else if habitat.season_day_count >= 270 && habitat.season_day_count < 360 {  // Fall
		return habitat.FoodLevel * 0.7, habitat.DangerLevel * 1.2
	}
	return habitat.FoodLevel * 1.0, habitat.DangerLevel * 1.0
}