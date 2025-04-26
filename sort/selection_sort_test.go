package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSortList(t *testing.T) {

	values := []int{33, 44, 55, 66, 78, 90, 100, 1, 10, 56}
	expected := []int{1, 10, 33, 44, 55, 56, 66, 78, 90, 100}

	result := SelectionSort(values)
	assert.Equal(t, expected, result)
}
