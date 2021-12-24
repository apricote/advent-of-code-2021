package main

var (
	programType [14]int = [14]int{1, 1, 1, 26, 1, 26, 1, 1, 1, 26, 26, 26, 26, 26}
	type2Number [14]int = [14]int{0, 0, 0, -14, 0, -7, 0, 0, 0, -7, -8, -7, -5, -10}
	type1Number [14]int = [14]int{15, 5, 6, 0, 9, 0, 14, 3, 1, 0, 0, 0, 0, 0}
)

func Backtrack(digits []int, direction bool) ([]int, bool) {
	if Reject(digits) {
		return []int{}, true
	}

	if Accept(digits) {
		return digits, false
	}

	next := First(digits, direction)
	for next != nil {
		solution, rejected := Backtrack(next, direction)
		if !rejected {
			return solution, false
		}

		next = Next(next, direction)
	}

	return []int{}, true
}

func Reject(digits []int) bool {
	z := 0
	for i, digit := range digits {
		switch programType[i] {
		case 1:
			z = 26*z + digit + type1Number[i]
		case 26:
			if z%26+type2Number[i] != digit {
				return true
			}

			z /= 26
		}
	}

	return false
}

func Accept(digits []int) bool {
	if len(digits) != 14 {
		return false
	}

	z := 0
	for i, digit := range digits {
		switch programType[i] {
		case 1:
			z = 26*z + digit + type1Number[i]
		case 26:
			if z%26+type2Number[i] != digit {
				return true
			}

			z /= 26
		}
	}

	return z == 0
}

func First(digits []int, direction bool) []int {
	if direction {
		digits = append(digits, 9)
	} else {
		digits = append(digits, 1)
	}
	return digits
}

func Next(digits []int, direction bool) []int {

	if direction {
		if digits[len(digits)-1] == 1 {
			return nil
		}

		digits[len(digits)-1] -= 1
	} else {
		if digits[len(digits)-1] == 9 {
			return nil
		}

		digits[len(digits)-1] += 1
	}

	return digits
}
