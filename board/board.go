package board

import (
	"bytes"
	"errors"
	"math"
	"strconv"
)

type Board struct {
	b     [9][9]int
	moves []pastMove
}

type pastMove struct {
	i, j, v int
}

func (b *Board) String() string {
	buf := make([]byte, 0, 90)
	for _, r := range b.b {
		for _, v := range r {
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func (b *Board) Get(i, j int) (int, error) {
	if i < 0 || j < 0 || i >= len(b.b) || j >= len(b.b) {
		return 0, errors.New("Out of bounds")
	}
	return b.b[i][j], nil
}

func Parse(puzzle []byte) (Board, error) {
	rows := bytes.Split(puzzle, []byte("\n"))
	b := Board{}
	for i, r := range rows {
		for j, v := range r {
			b.b[i][j], _ = strconv.Atoi(string(v))
		}
	}
	return b, nil
}

func (b *Board) PushMove(i, j, value int) bool {
	if !b.isLegalMove(i, j, value) {
		return false
	}
	b.moves = append(b.moves, pastMove{i, j, value})
	b.b[i][j] = value
	return true
}

func (b *Board) PopMove() (int, int, int) {
	if len(b.moves) == 0 {
		return 0, 0, 0
	}
	popped := b.moves[len(b.moves)-1]
	b.moves = b.moves[:len(b.moves)-1]
	b.b[popped.i][popped.j] = 0
	return popped.i, popped.j, popped.v
}

func (b *Board) isLegalMove(i, j, value int) bool {
	if i < 0 || i > 8 || j < 0 || j > 8 {
		return false
	}
	if b.b[i][j] != 0 {
		return false
	}
	if value < 1 || value > 9 {
		return false
	}
	return b.uniqueInRow(i, value) && b.uniqueInCol(j, value) && b.uniqueInBox(i, j, value)
}

func (b *Board) uniqueInRow(i, value int) bool {
	for _, v := range b.b[i] {
		if v == value {
			return false
		}
	}
	return true

}

func (b *Board) uniqueInCol(j, value int) bool {
	for _, row := range b.b {
		v := row[j]
		if v == value {
			return false
		}
	}
	return true
}

func box(i, j int) (int, int) {
	return i / 3 * 3, j / 3 * 3
}

func (b *Board) uniqueInBox(i, j, value int) bool {
	minX, minY := box(i, j)
	for _, row := range b.b[minX : minX+3] {
		for _, v := range row[minY : minY+3] {
			if v == value {
				return false
			}
		}
	}
	return true
}

func (b *Board) Solved() bool {
	for _, row := range b.b {
		for _, v := range row {
			if v == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) BestCell() (int, int) {
	i, j := 0, 0
	min := math.MaxInt32
	for y, row := range b.b {
		for z, val := range row {
			if val != 0 {
				continue
			}
			count := b.unfilledRow(y) + b.unfilledCol(z) + b.unfilledBox(y, z)
			if count < min {
				min = count
				i, j = y, z
			}
		}
	}
	return i, j
}

func contains(values []int, value int) (int, bool) {
	for x, v := range values {
		if v == value {
			return x, true
		}
	}
	return 0, false
}

func (b *Board) PossibleValues(i, j int) []int {
	possibilities := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range b.b[i] {
		if v == 0 {
			continue
		} else if x, c := contains(possibilities, v); c {
			possibilities = append(possibilities[:x], possibilities[x+1:]...)
		}
	}
	for _, row := range b.b {
		v := row[j]
		if v == 0 {
			continue
		} else if x, c := contains(possibilities, v); c {
			possibilities = append(possibilities[:x], possibilities[x+1:]...)
		}
	}
	minX, minY := box(i, j)
	for _, row := range b.b[minX : minX+3] {
		for _, v := range row[minY : minY+3] {
			if v == 0 {
				continue
			} else if x, c := contains(possibilities, v); c {
				possibilities = append(possibilities[:x], possibilities[x+1:]...)
			}

		}
	}
	return possibilities
}

func (b *Board) unfilledRow(i int) int {
	count := 0
	for _, v := range b.b[i] {
		if v == 0 {
			count++
		}
	}
	return count
}

func (b *Board) unfilledCol(j int) int {
	count := 0
	for _, row := range b.b {
		v := row[j]
		if v == 0 {
			count++
		}
	}
	return count
}

func (b *Board) unfilledBox(i, j int) int {
	minX, minY := box(i, j)
	count := 0
	for _, row := range b.b[minX : minX+3] {
		for _, v := range row[minY : minY+3] {
			if v == 0 {
				count++
			}
		}
	}
	return count
}
