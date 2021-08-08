package main

func main() {
	i := 0 // 1
L1: // 3
	for {
		if i == 0 {
			break L1 // 2
		}
	}
	println("OK") // 4
}
