package main

import (
	"strconv"
)

type Stack struct {
	letters []string
}

func (s *Stack) Push(letters string) {
	s.letters = append(s.letters, letters)
}

func (s *Stack) Pop() string {
	if len(s.letters) == 0 {
		panic("Stack is empty")
	}

	index := len(s.letters) - 1
	letters := s.letters[index]
	s.letters = s.letters[:index]

	return letters
}

// Task 3.a
func RobotTranslatorV2(cmd string) string {

	if cmd == "" {
		return ""
	}

	sliceOfString := convertToSliceOfString(cmd)

	stack := &Stack{}

	removePreviousLetterWithX(sliceOfString, stack)

	for _, val := range stack.letters {
		if val != "R" && val != "L" && val != "A" && val != "X" {
			return "Invalid Command"
		}
	}

	finalCommands := convertCommandToMovement(stack)

	return finalCommands

}

func convertToSliceOfString(cmd string) []string {
	var sliceOfString []string
	for i := range cmd {
		sliceOfString = append(sliceOfString, string(cmd[i]))
	}
	return sliceOfString
}

func removePreviousLetterWithX(sliceOfString []string, stack *Stack) {

	for _, val := range sliceOfString {
		if val == "X" {
			if len(stack.letters) == 0 {
				continue
			} else {
				stack.Pop()
			}
		} else {
			stack.Push(val)
		}
	}
}

func convertCommandToMovement(stack *Stack) string {

	if len(stack.letters) == 0 {
		return ""
	}

	currentLetter := stack.letters[0]
	sameCount := 1

	var result string

	for i := 1; i < len(stack.letters); i++ {
		if stack.letters[i] == currentLetter {
			sameCount += 1
		} else {
			switch currentLetter {
			case "R":
				if sameCount == 1 {
					result += "Move right " + strconv.Itoa(sameCount) + " time\n"
				} else {
					result += "Move right " + strconv.Itoa(sameCount) + " times\n"
				}
			case "L":
				if sameCount == 1 {
					result += "Move left " + strconv.Itoa(sameCount) + " time\n"
				} else {
					result += "Move left " + strconv.Itoa(sameCount) + " times\n"
				}
			case "A":
				if sameCount == 1 {
					result += "Move advance " + strconv.Itoa(sameCount) + " time\n"
				} else {
					result += "Move advance " + strconv.Itoa(sameCount) + " times\n"
				}
			}
			currentLetter = stack.letters[i]
			sameCount = 1
		}
	}

	var finalLetter string
	if currentLetter == "R" {
		finalLetter = "right"
	} else if currentLetter == "L" {
		finalLetter = "left"
	} else if currentLetter == "A" {
		finalLetter = "advance"
	}

	var amount string
	if sameCount > 1 {
		amount = "times"
	} else {
		amount = "time"
	}

	result += "Move " + finalLetter + " " + strconv.Itoa(sameCount) + " " + amount

	return result
}
