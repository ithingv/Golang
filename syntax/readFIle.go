package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type person []string

func main() {
	person := newPersonFromFile("person")
	// error
	//person := newPersonFromFile("person1")
	person.print()

}

func newName() person {
	names := person{}

	firstNames := []string{"Kim", "lee", "Park"}
	lastNames := []string{"hong", "hoon", "sung"}

	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			names = append(names, firstName+" "+lastName)
		}
	}
	return names
}

func (p person) print() {
	for i, name := range p {
		fmt.Println(i, name)
	}
}

func split(p person, splitSize int) (person, person) {
	return p[:splitSize], p[splitSize:]
}

func (p person) toString() string {
	return strings.Join([]string(p), ", ")
}

func (p person) toByte() []byte {
	return []byte(p.toString())
}

func (p person) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, p.toByte(), 0666)
}

func newPersonFromFile(filename string) person {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// 옵션1 - log the error and return a call to newName() 
		// 옵션2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ", ")
	return person(s)
}
