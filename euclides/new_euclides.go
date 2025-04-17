package euclides

func NewMdc(a, b int) int {
	aIsEven := a%2 == 0
	bIsEven := b%2 == 0

	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}

	if aIsEven && bIsEven {
		return 2 * NewMdc(a/2, b/2)
	}

	if !aIsEven && bIsEven {
		return NewMdc(a, b/2)
	}

	if aIsEven && !bIsEven {
		return NewMdc(a/2, b)
	}

	if a >= b {
		return NewMdc((a-b)/2, b)
	} else {
		return NewMdc((b-a)/2, a)
	}
}
