package allergies

import "slices"

const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var all = []uint{cats, pollen, chocolate, tomatoes, strawberries, shellfish, peanuts, eggs}

var point2Allergy = map[uint]string{
	eggs:         "eggs",
	peanuts:      "peanuts",
	shellfish:    "shellfish",
	strawberries: "strawberries",
	tomatoes:     "tomatoes",
	chocolate:    "chocolate",
	pollen:       "pollen",
	cats:         "cats",
}

func Allergies(allergies uint) []string {
	set := make(map[uint]struct{})
	index := 0
	for allergies > 0 && index < len(all) {
		if allergies&all[index] != 0 {
			set[all[index]] = struct{}{}
		}
		index++
	}
	res := make([]string, 0)
	for k := range set {
		res = append(res, point2Allergy[k])
	}
	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	set := Allergies(allergies)
	return slices.Contains(set, allergen)
}
