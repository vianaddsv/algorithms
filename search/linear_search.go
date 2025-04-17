package search

import "errors"

func Find(a []int, target int) (int, error) {
	for index, val := range a {
		if target == val {
			return index, nil
		}
	}
	return -1, ErrorCannotFoundTarget
}

var (
	ErrorCannotFoundTarget = errors.New("cannot found target")
)
