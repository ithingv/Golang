// 다양한 for문 사용법

package main

func main() {

	// 1부터 100까지 더하는 프로그램
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	// for 루프는 n이 100을 넘는지 체크하는 조건식만 사용 (while과 유사)
	n := 1
	for n < 100 {
		n *= 2
		// if n > 90{
		//	break
		//}
	}
	println(n)

	// 무한루프
	for {
		println("Infinite loop")
	}

	/* 3명의 이름을 갖는 문자열 배열에서 문자열 인덱스(0,1,2)와 해당 이름을
	   차례로 가져와 프린트하는 코드
	*/
	names := []string{"가", "나", "다"}
	for index, name := range names {
		println(index, name)
	}
}
