package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"seoul": {
		40.68433, -74.39967,
	},
	"busan": {
		40.68433, -74.39967,
	},
}

func main() {
	// m = make(map[string]Vertex)
	// m["seoul"] = Vertex{
	// 	40.0, -10.0,
	// }
	fmt.Println(m["seoul"])
	fmt.Println(m["busan"])

	delete(m, "busan")
	fmt.Println("The value:", m["busan"])

	v, ok := m["busan"]
	fmt.Println("The value:", v, "Present?", ok)
}
