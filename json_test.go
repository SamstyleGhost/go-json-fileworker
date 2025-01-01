package jsonfileworker

import (
	"fmt"
	"testing"

	"github.com/SamstyleGhost/go-json-fileworker/models"
)

const (
	usersFilePath = "data/users.json"
)

func TestGetUsers(t *testing.T) {
	var users []models.Users

	err := GetAllObjects(usersFilePath, &users)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(users, "", "  ")
	fmt.Print(string(b))
}

func TestGetUsersFromIndex(t *testing.T) {

	v, err := GetObjectFromIndex[models.Users](usersFilePath, 5)
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(b))
}
