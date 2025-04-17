package euclides

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectMdcForNewMdc(t *testing.T) {
	result := NewMdc(1035, 759)
	resultExpected := 69

	assert.Equal(t, result, resultExpected)
}
