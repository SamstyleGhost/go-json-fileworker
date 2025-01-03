package jsonfileworker

import (
	"fmt"
	"testing"

	"github.com/SamstyleGhost/go-json-fileworker/models"
)

const (
	usersFilePath = "data/users.json"
	trialFilePath = "data/trial.json"
	booksFilePath = "data/books.json"
)

func TestGetObjects(t *testing.T) {
	var users []models.Users

	err := GetAllObjects(usersFilePath, &users)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(users, "", "  ")
	fmt.Print(string(b))
}

func TestGetObjectsFromIndex(t *testing.T) {

	v, err := GetObjectFromIndex[models.Users](usersFilePath, 5)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(b))
}

func TestGetObjectsFromIndexForNonArrays(t *testing.T) {

	v, err := GetObjectFromIndex[models.Users](trialFilePath, 5)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(b))
}

func TestGetObjectsFromIndexForWrongType(t *testing.T) {

	v, err := GetObjectFromIndex[models.Users](booksFilePath, 5)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(b))
}

func TestAppendObjectsToArray(t *testing.T) {

	// Just pulled an object from users instead of creating a new one
	obj, err := GetObjectFromIndex[models.Users](usersFilePath, 0)
	if err != nil {
		t.Error(err)
	}

	err = AppendObjectToArray(usersFilePath, obj)
	if err != nil {
		t.Error(err)
	}
}
