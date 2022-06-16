package grid

func ValidDiagonalNeighbours(index, rowSize, rowCount int) []int {
	x, y := index%rowSize, index/rowSize
	neighbours := make([]int, 0, 8)

	top := y > 0
	bottom := y < rowCount-1

	left := x > 0
	right := x < rowSize-1

	if top {
		if left {
			neighbours = append(neighbours, index-rowSize-1)
		}
		neighbours = append(neighbours, index-rowSize)
		if right {
			neighbours = append(neighbours, index-rowSize+1)
		}
	}

	if left {
		neighbours = append(neighbours, index-1)
	}

	if right {
		neighbours = append(neighbours, index+1)
	}

	if bottom {
		if left {
			neighbours = append(neighbours, index+rowSize-1)
		}
		neighbours = append(neighbours, index+rowSize)
		if right {
			neighbours = append(neighbours, index+rowSize+1)
		}
	}

	return neighbours

}

func ValidNeighbours(index, rowSize, rowCount int) []int {
	x, y := index%rowSize, index/rowCount
	neighbours := make([]int, 0, 4)

	top := y > 0
	bottom := y < rowCount-1

	left := x > 0
	right := x < rowSize-1

	if top {
		neighbours = append(neighbours, index-rowSize)
	}

	if left {
		neighbours = append(neighbours, index-1)
	}

	if right {
		neighbours = append(neighbours, index+1)
	}

	if bottom {
		neighbours = append(neighbours, index+rowSize)
	}

	return neighbours
}

func intAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func Manhatten(x1, y1, x2, y2 int) int {
	return intAbs(x1-x2) + intAbs(y1-y2)
}

func ToXY(index, rowSize int) (int, int) {
	return index % rowSize, index / rowSize
}
