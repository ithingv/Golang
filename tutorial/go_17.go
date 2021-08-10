// type문을 사용한 함수 원형 정의
package main

// 원형 정의
type calculaotor func(int, int) int

// func calc(f func(int, int) int, a int, b int) int {
// calculator 원형 사용
func calc(f calculaotor, a int, b int) int {
	result := f(a, b)
	return result
}
