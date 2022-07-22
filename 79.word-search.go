package main

func Alg(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	if len(word) == 0 {
		return true
	}

	oneLineBoard := make([]byte, 0)

	for _, b := range board {
		oneLineBoard = append(oneLineBoard, b...)
	}

	excluded := make(map[int]bool)
	positions := make([]int, 0)

	for i, ch := range oneLineBoard {
		if ch == word[0] {
			positions = append(positions, i)
		}
	}

	for i, position := range positions {
		result := true
		excluded = make(map[int]bool)

		for j := range word {
			letterFound, nextPosition := findLetter(&excluded, oneLineBoard, position, word[j], len(board[0]))
			if !letterFound {
				if i == len(positions)-1 {
					return false
				}

				result = false
				break
			}

			position = nextPosition
		}

		if result {
			return result
		}
	}

	return len(positions) > 0
}

func findLetter(excluded *map[int]bool, board []byte, position int, letterToFind byte, rowLen int) (bool, int) {
	lettersPosition := -1

	defer func() {
		(*excluded)[lettersPosition] = true
	}()

	if _, ok := (*excluded)[position]; !ok && board[position] == letterToFind {
		lettersPosition = position
		return true, lettersPosition
	}

	// left
	if _, ok := (*excluded)[position-1]; !ok && position > 0 {
		if board[position-1] == letterToFind {
			lettersPosition = position - 1
			return true, lettersPosition
		}
	}

	// right
	if _, ok := (*excluded)[position+1]; !ok && position%rowLen < rowLen-1 {
		if board[position+1] == letterToFind {
			lettersPosition = position + 1
			return true, lettersPosition
		}
	}

	// down
	if _, ok := (*excluded)[position+rowLen]; !ok && position+rowLen < len(board) {
		if board[position+rowLen] == letterToFind {
			lettersPosition = position + rowLen
			return true, lettersPosition
		}
	}

	// up
	if _, ok := (*excluded)[position-rowLen]; !ok && position-rowLen >= 0 {
		if board[position-rowLen] == letterToFind {
			lettersPosition = position - rowLen
			return true, lettersPosition
		}
	}

	return false, lettersPosition
}
