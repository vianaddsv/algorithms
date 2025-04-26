package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSortAllElements(t *testing.T) {
	a := []int{33, 44, 55, 66, 78, 90, 100, 1, 10, 56}
	min := 1 //what is the best choice?
	expected := []int{1, 10, 33, 44, 55, 56, 66, 78, 90, 100}
	result := QuicksortSelection(a, min)

	assert.Equal(t, expected, result)

}
