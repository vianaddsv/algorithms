package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectIndexOfElement(t *testing.T) {
	result, err := Find([]int{1, 4, 6, 8, 10}, 10)
	resultExpected := 4

	assert.Equal(t, result, resultExpected)
	assert.Nil(t, err)

}

func TestShouldNotFound(t *testing.T) {
	result, err := Find([]int{1, 4, 6, 8, 10}, 100)
	resultExpected := -1

	assert.Equal(t, result, resultExpected)
	assert.ErrorIs(t, ErrorCannotFoundTarget, err)

}
