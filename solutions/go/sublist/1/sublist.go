package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	s1, s2 := len(l1), len(l2)
	if s1 == 0 && s2 == 0 {
		return RelationEqual
	}
	if s1 == 0 {
		return RelationSublist
	}
	if s2 == 0 {
		return RelationSuperlist
	}
	if s1 > s2 {
		relation := Sublist(l2, l1)
		if relation == RelationSublist {
			return RelationSuperlist
		}
		return relation
	}

	res := RelationUnequal
	for i := 0; i <= s2-s1; i++ {
		if l1[0] == l2[i] {
			if isSub(l1, l2[i:i+s1]) {
				res = RelationSublist
			}
		}
	}
	if s1 == s2 && res == RelationSublist {
		return RelationEqual
	}
	return res
}

func isSub(l1, l2 []int) bool {
	l, r := 0, 0
	for l < len(l1) && r < len(l2) {
		if l1[l] != l2[r] {
			return false
		}
		l++
		r++
	}
	return true
}
