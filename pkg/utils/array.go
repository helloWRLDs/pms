package utils

func ContainsInArray[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func MapToArray[K comparable, V any](m map[K]V) []V {
	arr := make([]V, len(m))
	i := 0
	for _, v := range m {
		arr[i] = v
		i++
	}
	return arr
}
