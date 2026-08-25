package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate * 0.01
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	f := float64(productionRate) * successRate * 0.01
	return int(f) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTen uint
	groupsOfTen = uint(carsCount / 10)

	var remain uint
	remain = uint(carsCount % 10)
	return groupsOfTen*95000 + remain*10000
}
