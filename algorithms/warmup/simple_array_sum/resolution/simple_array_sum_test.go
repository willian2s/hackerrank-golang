package simplearraysum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	resolution "github.com/willian2s/hackerrank-golang/algorithms/warmup/simple_array_sum/resolution"
)

func TestSimpleArraySum(t *testing.T) {
	assert.Equal(t, int32(3), resolution.SimpleArraySum([]int32{1, 2}))
	assert.Equal(t, int32(6), resolution.SimpleArraySum([]int32{1, 2, 3}))
	assert.Equal(t, int32(31), resolution.SimpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
}
