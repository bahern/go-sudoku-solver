package sudoku

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PuzzleSuite struct {
	suite.Suite
}

func TestPuzzleSuite(t *testing.T) {
	suite.Run(t, new(PuzzleSuite))
}

func (s PuzzleSuite) TestLoad_EmptyString() {
	// Given
	p := Puzzle{}
	r := io.Reader(strings.NewReader(""))

	// When
	err := p.Load(r)

	// Then
	s.NotNil(err)
	if err != nil {
		s.Equal("Invalid puzzle data.  Expected 9 rows, found 0 row(s)", err.Error())
	}
}

func (s PuzzleSuite) TestLoad_ShortRow() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/bad_row_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// When
	err = p.Load(file)

	// Then
	s.NotNil(err)
	if err != nil {
		s.Equal("Malformed row data: [6,x,x,1,9,5,x,x]", err.Error())
	}
}

func (s PuzzleSuite) TestLoad_LongRow() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/bad_row_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// When
	err = p.Load(file)

	// Then
	s.NotNil(err)
	if err != nil {
		s.Equal("Malformed row data: [6,x,x,1,9,5,x,x,x,x]", err.Error())
	}
}

func (s PuzzleSuite) TestLoad_InvalidValue() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/bad_row_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// When
	err = p.Load(file)

	// Then
	s.NotNil(err)
	if err != nil {
		s.Equal("Malformed row data, 10 is not between 1 and 9: [6,x,x,1,10,5,x,x,x]", err.Error())
	}
}

func (s PuzzleSuite) TestReduceRow() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/sample_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_ = p.Load(file)

	// Expect to reduce 3,4,5,8,9 from the set of potential values.
	expected := []bool{true, true, false, false, false, true, true, false, false}

	// When
	b := p.reduceRow(&p[0][0])

	// Then
	s.True(b)
	s.Equal(expected, p[0][0].potentialValues)
}

func (s PuzzleSuite) TestReduceColumn() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/sample_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_ = p.Load(file)

	// Expect to reduce 2,5,7,9 from the set of potential values.
	expected := []bool{true, false, true, true, false, true, false, true, false}

	// When
	b := p.reduceColumn(&p[0][0])

	// Then
	s.True(b)
	s.Equal(expected, p[0][0].potentialValues)
}

func (s PuzzleSuite) TestReduceGrid_1() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/sample_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_ = p.Load(file)

	// Expect to reduce 3,4,5,7,8,9 from the set of potential values.
	expected := []bool{true, true, false, false, false, true, false, false, false}

	// When
	b := p.reduceGrid(&p[0][0])

	// Then
	s.True(b)
	s.Equal(expected, p[0][0].potentialValues)
}

func (s PuzzleSuite) TestReduceGrid_2() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/sample_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_ = p.Load(file)

	// Expect to reduce 7 from the set of potential values.
	expected := []bool{true, true, true, true, true, true, false, true, true}

	// When
	b := p.reduceGrid(&p[4][5])

	// Then
	s.True(b)
	s.Equal(expected, p[4][5].potentialValues)
}

func (s PuzzleSuite) TestReduceGrid_3() {
	// Given
	p := Puzzle{}
	file, err := os.Open("../fixtures/sample_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_ = p.Load(file)

	// A cell with a value shouldn't touch this.
	expected := []bool{false, false, false, false, false, false, false, false, false}

	// When
	b := p.reduceGrid(&p[4][4])

	// Then
	s.False(b)
	s.Equal(expected, p[4][4].potentialValues)
}
