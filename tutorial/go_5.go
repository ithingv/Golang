switch v.(type) {
case int:
	println("int")
case string:
	println("string")
case bool:
	println("boo;")
default:
	println("unknown")
}