package allergies

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var res []string
	for allergen, score := range allergens {
		if allergies&score != 0 {
			res = append(res, allergen)
		}
	}
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	score, found := allergens[allergen]
	if !found {
		return false
	}
	return allergies&score != 0
}
