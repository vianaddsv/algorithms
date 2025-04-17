package search

import "errors"

func BinarySearch(a []int, target int) (int, error) {
	if len(a) == 0 {
		return -1, ErrorCannotFoundBinarySearchTarget
	}

	return binarySearch(a, target, 0, len(a)-1)

}

func binarySearch(a []int, target int, firstIndex int, lastIndex int) (int, error) {
	middleIndex := (firstIndex + lastIndex) / 2
	middleElement := a[middleIndex]

	if middleElement == target {
		return middleIndex, nil
	}
	if firstIndex == lastIndex {
		return -1, ErrorCannotFoundBinarySearchTarget
	}

	if middleElement > target {
		return binarySearch(a, target, firstIndex, middleIndex-1)
	}
	if middleElement < target {
		return binarySearch(a, target, middleIndex+1, lastIndex)
	}

	return -1, ErrorCannotFoundBinarySearchTarget
}

var (
	ErrorCannotFoundBinarySearchTarget = errors.New("cannot found target")
)
