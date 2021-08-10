// pass by value, pass by reference

package main

func main() {
	msg1 := "Hello"
	msg2 := "Hello"
	println("original msg1: ", msg1)
	say_pass_by_value(msg1)
	println("changed msg1: ", msg1)
	println("original msg2: ", msg2)
	say_pass_by_reference(&msg2)
	println("changed msg2: ", msg2)
}

func say_pass_by_value(msg string) {
	msg = "world"
}

func say_pass_by_reference(msg *string) {
	*msg = "world"
}
