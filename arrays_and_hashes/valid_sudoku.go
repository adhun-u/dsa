package arraysandhashes

func containsInMap(hashMap map[byte]int, numToCheck byte) bool {
	count, _ := hashMap[numToCheck]
	return count > 1
}

func matrixKey(i int, j int) int {

	if i >= 0 && i <= 2 {
		if j >= 0 && j <= 2 {
			return 0
		} else if j >= 3 && j <= 5 {
			return 1
		} else {
			return 2
		}
	} else if i >= 3 && i <= 5 {

		if j >= 0 && j <= 2 {
			return 3
		} else if j >= 3 && j <= 5 {
			return 4
		} else {
			return 5
		}
	} else {
		if j >= 0 && j <= 2 {
			return 6
		} else if j >= 3 && j <= 5 {
			return 7
		} else {
			return 8
		}
	}
}

func IsValidSudoku(board [][]byte) bool {

	columns := make(map[int]map[byte]int)
	matrices := make(map[int]map[byte]int)
	rows := make(map[int]map[byte]int)

	for i := range board {

		rowInBoard := board[i]
		rowMap, rExists := rows[i]

		if !rExists {
			rowMap = map[byte]int{}
		}

		for j := range rowInBoard {
			//Row checking and adding
			rowElem := rowInBoard[j]
			if rowElem != '.' {
				rowMap[rowElem]++

				contains := containsInMap(rowMap, rowElem)
				if contains {
					return false
				}
			}

			rows[i] = rowMap

			//3x3 Matrix checking and adding
			index := matrixKey(i, j)
			matrix, mExists := matrices[index]
			if !mExists {
				matrix = map[byte]int{}
			}
			if rowElem != '.' {
				matrix[rowElem]++

				contains := containsInMap(matrix, rowElem)
				if contains {
					return false
				}
			}
			matrices[index] = matrix

			//Column checking and adding
			columnMap, cExists := columns[i]
			if !cExists {
				columnMap = map[byte]int{}
			}
			columnElem := board[j][i]

			if columnElem != '.' {

				columnMap[columnElem]++

				contains := containsInMap(columnMap, columnElem)
				if contains {
					return false
				}
			}
			columns[i] = columnMap

		}

	}
	return true
}
