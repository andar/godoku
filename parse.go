package main

import (
	"bytes"
	"errors"
	"strconv"
)

type Board struct {
	b [9][9]int
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
	return b.isLegalMove(i, j, value)
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
