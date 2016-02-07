package board

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	puzzle, err := ioutil.ReadFile("../example.txt")
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
	puzzle, err := ioutil.ReadFile("../example.txt")
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

var moveSequence = []move{
	{8, 8, 2, true},
	{5, 4, 5, true},
	{1, 4, 7, true},
	{0, 3, 5, true},
}

func TestBoardPopMove(t *testing.T) {
	puzzle, err := ioutil.ReadFile("../example.txt")
	if err != nil {
		t.Fatal(err)
	}
	board, err := Parse(puzzle)
	if err != nil {
		t.Fatal(err)
	}
	for _, move := range moveSequence {
		success := board.PushMove(move.i, move.j, move.v)
		if success != move.success {
			t.Fatalf("Expected success of (%d, %d, %d) to be %t", move.i, move.j, move.v, move.success)
		}
		v, _ := board.Get(move.i, move.j)
		if v != move.v {
			t.Fatalf("Expected (%d, %d) to equal %d, got %d", move.i, move.j, move.v, v)
		}
	}
	for x := len(moveSequence) - 1; x >= 0; x-- {
		undoneMove := moveSequence[x]
		i, j, v := board.PopMove()
		if i != undoneMove.i || j != undoneMove.j || v != undoneMove.v {
			t.Fatalf("Expected popped move to equal (%d, %d, %d); got (%d, %d, %d)",
				undoneMove.i, undoneMove.j, undoneMove.v, i, j, v)
		}
	}
}

func TestBestCell(t *testing.T) {
	puzzle, err := ioutil.ReadFile("../example.txt")
	if err != nil {
		t.Fatal(err)
	}
	board, err := Parse(puzzle)
	if err != nil {
		t.Fatal(err)
	}
	i, j := board.BestCell()
	v, _ := board.Get(i, j)
	if v != 0 {
		t.Fatalf("Invalid starting cell (%d, %d)", i, j)
	}
	success := board.PushMove(i, j, 2)
	if !success {
		t.Fatal("Expected success")
	}
	newi, newj := board.BestCell()
	if newi == i && newj == j {
		t.Fatalf("Got same cell as before: new (%d, %d); old (%d, %d)", newi, newj, i, j)
	}
	board.PopMove()
	newi, newj = board.BestCell()
	if newi != i && newj != j {
		t.Fatal("Expected same move as before")
	}
}

func TestPossibleValues(t *testing.T) {
	puzzle, err := ioutil.ReadFile("../example.txt")
	if err != nil {
		t.Fatal(err)
	}
	board, err := Parse(puzzle)
	if err != nil {
		t.Fatal(err)
	}
	i, j := board.BestCell()
	v, _ := board.Get(i, j)
	if v != 0 {
		t.Fatalf("Invalid starting cell (%d, %d)", i, j)
	}
	possibleValues := board.PossibleValues(i, j)
	for _, v := range possibleValues {
		success := board.PushMove(i, j, v)
		if !success {
			t.Fatalf("Expected success for move (%d, %d, %d)", i, j, v)
		}
		board.PopMove()
	}
}
