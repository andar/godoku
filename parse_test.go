package main

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	puzzle, err := ioutil.ReadFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	board, err := Parse(puzzle)
	if err != nil {
		t.Fatal(err)
	}
	outOfBounds := [][]int{{-1, -1}, {-1, 0}, {0, -1}, {10, 10}, {0, 10}, {10, 0}}
	for _, v := range outOfBounds {
		_, err = board.Get(v[0], v[1])
		if err == nil {
			t.Fatal("Expected error")
		}
	}
	v, err := board.Get(1, 1)
	if err != nil {
		t.Fatal("Unexpeted error")
	}
	if v != 5 {
		t.Fatalf("Board position is incorrect: got %d, wanted %d", v, 5)
	}
}

type move struct {
	i, j, v int
	success bool
}

var exampleMoves = []move{
	{1, 1, 8, false},
	{0, 0, 10, false},
	{0, 0, -1, false},
	{0, 0, 0, false},
	{-1, -1, 1, false},
	{-1, 0, 1, false},
	{0, -1, 1, false},
	{9, 9, 1, false},
	{9, 0, 1, false},
	{0, 9, 1, false},
	{0, 0, 6, false}, // value exists in row
	{0, 0, 7, false}, // value exists in column
	{0, 0, 5, false}, // value exists in box
	{8, 8, 9, false}, // value exists in box
	{8, 8, 2, true},
	{5, 4, 5, true},
	{5, 4, 7, false},
	{5, 4, 9, false},
	{5, 4, 1, false},
	{1, 4, 7, true},
	{1, 4, 6, false},
	{1, 4, 2, false},
	{1, 4, 5, false},
	{0, 3, 5, true},
}

func TestBoardPushMove(t *testing.T) {
	puzzle, err := ioutil.ReadFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	for _, example := range exampleMoves {
		board, err := Parse(puzzle)
		if err != nil {
			t.Fatal(err)
		}

		success := board.PushMove(example.i, example.j, example.v)
		if success != example.success {
			t.Fatalf("Expected success of (%d, %d, %d) to be %t", example.i, example.j, example.v, example.success)
		}
	}
}
