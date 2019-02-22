package main

func findIndex(i int, j int) int {
	if i < 3 {
		if j < 3 {
			return 0
		} else if j < 6 {
			return 1
		} else {
			return 2
		}
	} else if i < 6 {
		if j < 3 {
			return 3
		} else if j < 6 {
			return 4
		} else {
			return 5
		}
	} else {
		if j < 3 {
			return 6
		} else if j < 6 {
			return 7
		} else {
			return 8
		}
	}
}

func isValidSudoku(board [][]byte) bool {
	var row []map[byte]byte
	var box []map[byte]byte
	i := 0
	for ; i < 9; i++ {
		j := 0
		setI := make(map[byte]byte)
		for ; j < 9; j++ {
			var setJ map[byte]byte
			if len(row) < j+1 {
				setJ = make(map[byte]byte)
				row = append(row, setJ)
			} else {
				setJ = row[j]
			}
			indexB := findIndex(i, j)
			var setB map[byte]byte
			if len(box) < indexB+1 {
				setB = make(map[byte]byte)
				box = append(box, setB)
			} else {
				setB = box[indexB]
			}
			value := board[i][j]
			if value == '.' {
				continue
			} else if _, ok := setI[value]; ok {
				return false
			} else {
				setI[value] = value
				if _, ok := setJ[value]; ok {
					return false
				} else {
					setJ[value] = value
					if _, ok := setB[value]; ok {
						return false
					} else {
						setB[value] = value
					}
				}
			}
		}
	}
	return true
}

func main() {
	nums := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	println(isValidSudoku(nums))
}
