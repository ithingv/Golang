package main

func main() {
	a := 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for 루프 시작으로
		}
		a++
		if a > 10 {
			break // 루프 빠져나옴
		}
	}
	if a == 11 {
		goto END // goto 사용예
	}
	println(a)
END:
	println("END")
}
