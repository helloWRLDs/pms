package utils

func If[T comparable](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}
