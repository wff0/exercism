package lasagnamaster

func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, l := range layers {
		if l == "noodles" {
			noodles += 50
		} else if l == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	tmp := make([]float64, len(quantities))
	for i := range quantities {
		tmp[i] = quantities[i] / 2 * float64(portions)
	}
	return tmp
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
