package spaceage

type Planet string

func (p Planet) Rate() float64 {
	switch p {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Earth":
		return 1
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	}
	return 0
}

const OneYearSec = 31557600

func Age(seconds float64, planet Planet) float64 {
	if planet.Rate() == 0 {
		return -1
	}
	return seconds / planet.Rate() / OneYearSec
}
