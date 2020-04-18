package leetcode

func GameOfLife(board [][]int) (out [][]int) {
	out = make([][]int, len(board))
	// Check the values around each point
	xLimit := len(board) - 1
	for x, _ := range board {
		yLimit := len(board[x]) - 1
		out[x] = make([]int, len(board[x]))
		for y, _ := range board[x] {
			liveCount := count(board, x, y, xLimit, yLimit)
			//fmt.Printf("change log ------- %d, %d, old: %d, liveCount: %d, new: %d\n", x, y, board[x][y], liveCount, change(liveCount, board[x][y]))
			out[x][y] = change(liveCount, board[x][y])
		}
	}
	copy(board, out)
	return
}

func count(in [][]int, x, y, xl, yl int) (liveCount int) {
	liveCount = 0
	for i := -1; i <= 1; i++ {
		px := x + i
		if px < 0 || px > xl {
			continue
		}
		for j := -1; j <= 1; j++ {
			py := y + j
			if py < 0 || py > yl {
				continue
			}
			if px == x && py == y {
				continue
			}
			if in[px][py] == 1 {
				liveCount++
			}
		}
	}
	return
}

func change(lc, value int) (out int) {
	if value == 1 {
		// live now
		if lc < 2 {
			return 0
		}
		if lc == 2 || lc == 3 {
			return 1
		}
		if lc > 3 {
			return 0
		}
	} else if value == 0 {
		if lc == 3 {
			return 1
		}
	}
	return value
}
