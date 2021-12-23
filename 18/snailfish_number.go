package main

import "strconv"

type SnailfishNumber interface{}

type SnailfishNumberPair [2]SnailfishNumber
type SnailfishNumberRegular int

//func main() {
//	number1 := SnailfishNumberPair{1, 2}
//	number2 := SnailfishNumberPair{number1, 3}
//}

func ParseSnailfishNumber(input string) (SnailfishNumber, error) {
	state := "none"
	var root *SnailfishNumberPair
	current := root

	for _, token := range input {
		switch token {
		case '[':

			if root == nil {
				root = &SnailfishNumberPair{}
				current = root
			} else {
				newPair := &SnailfishNumberPair{}
				switch state {
				case "open_pair":
					current[0] = newPair
				case "middle_pair":
					current[1] = newPair
				}
				current = newPair
			}

			state = "open_pair"
		case ']':
			state = "close_pair"
		case ',':
			state = "middle_pair"
		default:
			regular, err := strconv.Atoi(string(token))
			if err != nil {
				return SnailfishNumberPair{}, err
			}

			switch state {
			case "open_pair":
				current[0] = regular
			case "middle_pair":
				current[1] = regular
			}
		}
	}

	return *root, nil
}
