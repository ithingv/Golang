// 가변인자함수

package main

func main() {
	say("This", "is", "a", "book")
}

func say(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}
