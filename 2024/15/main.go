package main

import (
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/15/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	split := strings.Index(data, "\n\n")
	max_cols := strings.Index(data[:split], "\n")
	bytes := []byte(data)
	grid := bytes[0:split]
	robot_index := strings.Index(data, "@")
	result := 0
	for _, dir := range bytes[split:] {
		if dir == '\n' {
			continue
		}
		robot_index = tryMoveRobot(grid, max_cols, robot_index, dir)
	}
	for i, c := range grid {
		switch c {
		case '\n':
			continue
		case 'O':
			row, col := indexToCoord(max_cols, i)
			result += (row * 100) + col
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	split := strings.Index(data, "\n\n")
	max_rows := strings.Count(data[:split], "\n") + 1
	max_cols := strings.Index(data[:split], "\n") * 2
	bytes := []byte(data)
	grid := make([]byte, 0, (max_cols+1)*max_rows)
	robot_index := 0
	result := 0
	for _, c := range bytes[0:split] {
		switch c {
		case '#', '.':
			grid = append(grid, c, c)
		case 'O':
			grid = append(grid, '[', ']')
		case '@':
			robot_index = len(grid)
			grid = append(grid, '@', '.')
		case '\n':
			grid = append(grid, c)
		}
	}
	for _, dir := range bytes[split:] {
		if dir == '\n' {
			continue
		}
		robot_index = tryMoveRobot(grid, max_cols, robot_index, dir)
	}
	for i, c := range grid {
		switch c {
		case '\n':
			continue
		case '[':
			row, col := indexToCoord(max_cols, i)
			result += (row * 100) + col
		}
	}
	return result, nil
}

func indexToCoord(max_cols int, index int) (int, int) {
	row := index / (max_cols + 1)
	return row, index - row - (row * max_cols)
}

func coordToIndex(max_cols int, row int, col int) int {
	return (row * max_cols) + row + col
}

func move(row int, col int, dir byte) (int, int) {
	switch dir {
	case '^':
		row -= 1
	case '>':
		col += 1
	case 'v':
		row += 1
	case '<':
		col -= 1
	}
	return row, col
}

func tryMoveRobot(grid []byte, max_cols int, cur_idx int, dir byte) int {
	cur_row, cur_col := indexToCoord(max_cols, cur_idx)
	for {
		new_row, new_col := move(cur_row, cur_col, dir)
		new_idx := coordToIndex(max_cols, new_row, new_col)
		switch grid[new_idx] {
		case '.':
			grid[new_idx] = '@'
			grid[cur_idx] = '.'
			return new_idx
		case '#':
			return cur_idx
		case 'O':
			if tryMoveBox(grid, max_cols, new_idx, dir) {
				grid[new_idx] = '@'
				grid[cur_idx] = '.'
				return new_idx
			}
			return cur_idx
		case '[', ']':
			if tryMoveBigBox(grid, max_cols, new_idx, dir) {
				grid[new_idx] = '@'
				grid[cur_idx] = '.'
				return new_idx
			}
			return cur_idx
		}
	}
}

func tryMoveBox(grid []byte, max_cols int, cur_idx int, dir byte) bool {
	cur_row, cur_col := indexToCoord(max_cols, cur_idx)
	new_row, new_col := cur_row, cur_col
	for {
		new_row, new_col = move(new_row, new_col, dir)
		new_idx := coordToIndex(max_cols, new_row, new_col)
		switch grid[new_idx] {
		case '#':
			return false
		case '.':
			grid[new_idx] = 'O'
			grid[cur_idx] = '.'
			return true
		}
	}
}

func tryMoveBigBox(grid []byte, max_cols int, cur_idx int, dir byte) bool {
	if grid[cur_idx] == '.' {
		return true
	}
	cur_row, cur_col := indexToCoord(max_cols, cur_idx)
	if dir == '<' || dir == '>' {
		new_row, new_col := cur_row, cur_col
		if dir == '<' {
			new_col -= 1
		} else {
			new_col += 1
		}
		for {
			new_row, new_col = move(new_row, new_col, dir)
			new_idx := coordToIndex(max_cols, new_row, new_col)
			if grid[new_idx] == '#' {
				return false
			}
			if grid[new_idx] == '.' || tryMoveBigBox(grid, max_cols, new_idx, dir) {
				if dir == '<' {
					grid[new_idx] = '['
					grid[new_idx+1] = ']'
				} else {
					grid[new_idx] = ']'
					grid[new_idx-1] = '['
				}
				grid[cur_idx] = '.'
				return true
			}
			return false
		}
	} else {
		cur_idx_l, cur_idx_r := cur_idx, cur_idx
		new_row_l, new_col_l, new_row_r, new_col_r := cur_row, cur_col, cur_row, cur_col
		if grid[cur_idx] == '[' {
			new_col_r += 1
			cur_idx_r += 1
		} else {
			new_col_l -= 1
			cur_idx_l -= 1
		}
		for {
			new_row_l, new_col_l = move(new_row_l, new_col_l, dir)
			new_idx_l := coordToIndex(max_cols, new_row_l, new_col_l)
			new_row_r, new_col_r = move(new_row_r, new_col_r, dir)
			new_idx_r := coordToIndex(max_cols, new_row_r, new_col_r)
			if grid[new_idx_l] == '#' || grid[new_idx_r] == '#' {
				return false
			}
			if grid[new_idx_l] == '.' && grid[new_idx_r] == '.' {
				grid[new_idx_l] = '['
				grid[new_idx_r] = ']'
				grid[cur_idx_l] = '.'
				grid[cur_idx_r] = '.'
				return true
			}
			if grid[new_idx_l] == '[' && grid[new_idx_r] == ']' {
				if tryMoveBigBox(grid, max_cols, new_idx_l, dir) {
					grid[new_idx_l] = '['
					grid[new_idx_r] = ']'
					grid[cur_idx_l] = '.'
					grid[cur_idx_r] = '.'
					return true
				}
				return false
			}
			if canMoveBigBox(grid, max_cols, new_idx_l, dir) && canMoveBigBox(grid, max_cols, new_idx_r, dir) {
				_ = tryMoveBigBox(grid, max_cols, new_idx_l, dir)
				_ = tryMoveBigBox(grid, max_cols, new_idx_r, dir)
				grid[new_idx_l] = '['
				grid[new_idx_r] = ']'
				grid[cur_idx_l] = '.'
				grid[cur_idx_r] = '.'
				return true
			}
			return false
		}
	}
}

func canMoveBigBox(grid []byte, max_cols int, cur_idx int, dir byte) bool {
	if grid[cur_idx] == '.' {
		return true
	}
	cur_row, cur_col := indexToCoord(max_cols, cur_idx)
	cur_idx_l, cur_idx_r := cur_idx, cur_idx
	new_row_l, new_col_l, new_row_r, new_col_r := cur_row, cur_col, cur_row, cur_col
	if grid[cur_idx] == '[' {
		new_col_r += 1
		cur_idx_r += 1
	} else {
		new_col_l -= 1
		cur_idx_l -= 1
	}
	new_row_l, new_col_l = move(new_row_l, new_col_l, dir)
	new_idx_l := coordToIndex(max_cols, new_row_l, new_col_l)
	new_row_r, new_col_r = move(new_row_r, new_col_r, dir)
	new_idx_r := coordToIndex(max_cols, new_row_r, new_col_r)
	if grid[new_idx_l] == '#' || grid[new_idx_r] == '#' {
		return false
	}
	if grid[new_idx_l] == '.' && grid[new_idx_r] == '.' {
		return true
	}
	if grid[new_idx_l] == '[' && grid[new_idx_r] == ']' {
		return canMoveBigBox(grid, max_cols, new_idx_l, dir)

	}
	return canMoveBigBox(grid, max_cols, new_idx_l, dir) && canMoveBigBox(grid, max_cols, new_idx_r, dir)
}
