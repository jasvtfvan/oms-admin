package utils

import "regexp"

// 切片（/数组）包含方法
func Contains[T comparable](slice []T, item T) bool {
	m := make(map[T]struct{})
	for _, val := range slice {
		m[val] = struct{}{}
	}
	_, exists := m[item]
	return exists
}

// 检查密码是否满足条件
func IsValidPassword(password string) bool {
	// 检查长度
	if len(password) < 8 {
		return false
	}
	// 检查是否包含小写字母
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	// 检查是否包含大写字母
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	// 检查是否包含数字
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false
	}
	return true
}
