package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type person []string

func main() {
	person := newName()
	//fmt.Println(person.toByte())
	person.saveToFile("person")

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
