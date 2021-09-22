package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

func (p person) shuffle() {

	source := rand.NewSource(time.Now().UnixNano()) // seed를 int64로 생성
	r := rand.New(source)

	for i := range p {
		newPos := r.Intn(len(p) - 1) // seed가 고정돼있음, type Source
		p[i], p[newPos] = p[newPos], p[i]
	}
}
