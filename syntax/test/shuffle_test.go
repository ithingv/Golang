package main

import (
	"os"
	"testing"
)

func TestNewName(t *testing.T) {
	names := newName()

	if len(names) != 9 {
		t.Errorf("Expected Names length of 9, but got %v", len(names))
	}

	if names[0] != "Kim hong" {
		t.Errorf("Expected Name Kim hong, but got %v", names[0])
	}

	if names[len(names)-1] != "Park sung" {
		t.Errorf("Expected Name Park sung, but got %v", names[len(names)-1])
	}
}

// Testing FileIO

func TestSaveToPersonAndNewPersonFromFile(t *testing.T) {
	os.Remove("_persontesting")

	person := newName()
	person.saveToFile("_persontesting")

	loadedPerson := newPersonFromFile("_persontesting")

	if len(loadedPerson) != 9 {
		t.Errorf("Expected 9 persons but %v", len(loadedPerson))
	}

	os.Remove("_persontesting")
}
