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
