/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
	"unicode"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	symbols := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	var res *int = nil

	var leftNum *int = nil
	var sign *string = nil

	middlemanResults := make([]int, 0)

	multiplicationIndexes := make(map[int]bool)

	// making all multiplication('*') and dividing('/') operations
	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			buf, nextIndex := getNum(s, i)

			if leftNum == nil && sign == nil {
				leftNum = &buf
				i = nextIndex

				continue
			}

			if sign != nil {
				// for (3 * 2 / 1) cases
				if leftNum == nil && *sign != "-" && *sign != "+" {
					leftNum = &middlemanResults[len(middlemanResults)-1]
					middlemanResults = middlemanResults[:len(middlemanResults)-1]
				}

				switch *sign {
				case "*":
					middlemanResults = append(middlemanResults, *leftNum*buf)

					multiplicationIndexes[i] = true
					multiplicationIndexes[getPrevNumStartIndex(s, i-1)] = true
					sign = nil
					leftNum = nil
				case "/":
					middlemanResults = append(middlemanResults, *leftNum/buf)

					multiplicationIndexes[i] = true
					multiplicationIndexes[getPrevNumStartIndex(s, i-1)] = true

					sign = nil
					leftNum = nil
				default:
					leftNum = &buf
				}
			}

			i = nextIndex
		}

		if i < len(s) {
			if _, ok := symbols[string(s[i])]; ok {
				bufStr := string(s[i])
				sign = &bufStr
			}

			i++
		}
	}

	leftNum = nil
	sign = nil
	var rightNum *int = nil
	// making add ('+') and substraction ('-') operations
	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			buf, nextIndex := getNum(s, i)

			// if the first digit is negative(-1 - 2 - 3)
			if i == 1 && sign != nil {
				buf *= -1
				sign = nil
			}

			if leftNum == nil {
				leftNum = &buf
				i = nextIndex

				continue
			}

			if sign != nil { // then it is the right number
				if _, ok := multiplicationIndexes[i]; ok {
					rightNum = &middlemanResults[0]
					if len(middlemanResults) > 1 {
						middlemanResults = middlemanResults[1:]
					} else {
						middlemanResults = make([]int, 0)
					}

					nextIndex = findTheEndOfMultiplication(s, i)
				} else {
					rightNum = &buf
				}
			}

			if sign != nil {
				switch *sign {
				case "+":
					if res == nil {
						tmp := 0
						res = &tmp
					}
					*res = *leftNum + *rightNum

					rightNum = nil
					*leftNum = *res
				case "-":
					if res == nil {
						tmp := 0
						res = &tmp
					}
					*res = *leftNum - *rightNum

					rightNum = nil
					*leftNum = *res
				default:
					*leftNum = *rightNum
					rightNum = nil
				}
			}

			i = nextIndex
		}

		if i < len(s) {
			if _, ok := symbols[string(s[i])]; ok {

				bufStr := string(s[i])
				sign = &bufStr
			}

			i++
		}
	}

	if res == nil {
		if len(middlemanResults) == 0 {
			if leftNum != nil {
				return *leftNum
			}

			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			return num
		}

		return middlemanResults[0]
	}

	return *res
}

func getNum(s string, index int) (number, whereFinish int) {
	buf := ""

	for ; index < len(s); index++ {
		r := rune(s[index])

		if unicode.IsDigit(r) {
			buf += string(r)
		} else {
			break
		}
	}

	num, err := strconv.Atoi(buf)
	if err != nil {
		panic(err.Error())
	}

	return num, index
}

func getPrevNumStartIndex(s string, index int) int {
	for i := index - 1; i > 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			continue
		}

		return i + 1
	}

	return 0
}

func findTheEndOfMultiplication(s string, startIndex int) int {
	for i := startIndex; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" {
			return i
		}
	}

	return len(s)
}
// @lc code=end
