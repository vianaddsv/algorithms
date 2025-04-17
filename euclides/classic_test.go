package euclides

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectMdc(t *testing.T) {
	result := Mdc(1035, 759)
	resultExpected := 69

	assert.Equal(t, result, resultExpected)
}
