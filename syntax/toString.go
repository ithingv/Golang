package main

import (
	"fmt"
	"strings"
)

type person []string

func main() {
	person := newName()
	fmt.Println(person.toString())

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
