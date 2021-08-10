// 슬라이스
package main

import "fmt"

func main() {
	var a []int        // 슬라이스 변수 선언, 크기 지정 하지않음
	a = []int{1, 2, 3} // 슬라이스에 리터럴값 지정
	a[1] = 10
	println(a)     // [3/3]0xc00003a760 메모리 주소출력
	fmt.Println(a) // [1, 10, 3]
}
