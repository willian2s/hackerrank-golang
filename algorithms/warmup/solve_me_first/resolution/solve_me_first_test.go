package solvemefirst_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	solvemefirst "github.com/willian2s/hackerrank-golang/algorithms/warmup/solve_me_first/resolution"
)

func TestSolveMeFirst(t *testing.T) {
	expected := uint32(3)
	assert.Equal(t, expected, solvemefirst.SolveMeFirst(1, 2))
}
