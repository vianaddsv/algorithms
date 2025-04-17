package euclides

func Mdc(a, b int) int {
	if b == 0 {
		return a
	}
	return Mdc(b, a%b)
}
