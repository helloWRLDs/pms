package utils

func Ptr[T comparable](value T) *T {
	var null T
	if value == null {
		return nil
	}
	return &value
}

func Value[T comparable](ptr *T) T {
	var null T
	if ptr == nil {
		return null
	}
	return *ptr
}
