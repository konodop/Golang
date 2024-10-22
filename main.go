package main

import (
	"fmt"
)

func Calc(expression string) (float64, error) {
	var left, right int = 0, 0
	var did int = 0
	for _, g := range expression {
		if did == 0 {
			if g == '1' {
				left *= 10
				left += 1
			} else if g == '2' {
				left *= 10
				left += 2
			} else if g == '3' {
				left *= 10
				left += 3
			} else if g == '4' {
				left *= 10
				left += 4
			} else if g == '5' {
				left *= 10
				left += 5
			} else if g == '6' {
				left *= 10
				left += 6
			} else if g == '7' {
				left *= 10
				left += 7
			} else if g == '8' {
				left *= 10
				left += 8
			} else if g == '9' {
				left *= 10
				left += 9
			} else if g == '0' {
				left *= 10
			} else if g == '+' {
				did = 1
			} else if g == '-' {
				did = 2
			} else if g == '*' {
				did = 3
			} else if g == '/' {
				did = 4
			}
		} else {
			if g == '1' {
				right *= 10
				right += 1
			} else if g == '2' {
				right *= 10
				right += 2
			} else if g == '3' {
				right *= 10
				right += 3
			} else if g == '4' {
				right *= 10
				right += 4
			} else if g == '5' {
				right *= 10
				right += 5
			} else if g == '6' {
				right *= 10
				right += 6
			} else if g == '7' {
				right *= 10
				right += 7
			} else if g == '8' {
				right *= 10
				right += 8
			} else if g == '9' {
				right *= 10
				right += 9
			} else if g == '0' {
				right *= 10
			}
		}
	}
	if did == 1 {
		return float64(left) + float64(right), nil
	} else if did == 2 {
		return float64(left) - float64(right), nil
	} else if did == 3 {
		return float64(left) * float64(right), nil
	} else {
		if right == 0 {
			return 0.0, fmt.Errorf("division by zero")
		}
		return float64(left) / float64(right), nil
	}
}

func main() {
	fmt.Println(Calc("4123 + 1234"))
}
