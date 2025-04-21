package utils

func Ptr[T comparable](value T) *T {
	var null T
	if value == null {
		return nil
	}
	return &value
}
