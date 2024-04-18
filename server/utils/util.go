package utils

func Contains[T comparable](slice []T, item T) bool {
	m := make(map[T]struct{})
	for _, val := range slice {
		m[val] = struct{}{}
	}
	_, exists := m[item]
	return exists
}
