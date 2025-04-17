package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFoundForLastItem(t *testing.T) {
	result, err := BinarySearch([]int{1, 4, 6, 8, 10}, 10)
	resultExpected := 4

	assert.Equal(t, resultExpected, result)
	assert.Nil(t, err)

}

func TestShouldFoundFirstItem(t *testing.T) {
	result, err := BinarySearch([]int{1, 4, 6, 8, 10}, 1)
	resultExpected := 0

	assert.Equal(t, result, resultExpected)
	assert.Nil(t, err)

}

func TestShouldFoundItem(t *testing.T) {
	result, err := BinarySearch([]int{1, 4, 6, 8, 10, 12, 20, 90, 100, 200, 450, 451}, 451)
	resultExpected := 11

	assert.Equal(t, result, resultExpected)
	assert.Nil(t, err)

}

func TestShouldNotFoundItem(t *testing.T) {
	result, err := BinarySearch([]int{1, 4, 6, 8, 10, 12, 20, 90, 100, 200, 450, 451}, 1000)
	resultExpected := -1

	assert.Equal(t, result, resultExpected)
	assert.ErrorIs(t, err, ErrorCannotFoundBinarySearchTarget)

}
