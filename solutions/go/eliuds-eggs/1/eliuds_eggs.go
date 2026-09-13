package eliudseggs

import "math/bits"

func EggCount(displayValue int) int {
	count64 := bits.OnesCount64(uint64(displayValue))
	return count64
}
