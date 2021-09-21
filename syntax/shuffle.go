package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type person []string

func main() {
	persons := newName()
	//fmt.Println(person.toByte())
	//persons.saveToFile("person")
	persons.shuffle()
	persons.print()
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

func (p person) shuffle() {

	source := rand.NewSource(time.Now().UnixNano()) // seed를 int64로 생성
	r := rand.New(source)

	for i := range p {
		newPos := r.Intn(len(p) - 1) // seed가 고정돼있음, type Source
		p[i], p[newPos] = p[newPos], p[i]
	}
}

func swap(p person, idx1 *int, idx2 *int) {

}
