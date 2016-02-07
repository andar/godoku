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
