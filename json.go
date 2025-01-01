package jsonfileworker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)

/*
	Add the following function:
	UpdateObject(pathToFile string, v interface{}, conditions, whatToUpdate)

	The conditions can be the following:
	- Index of element to be updated
	- Some equals, greaterthan, lessthan conditions (could use an enum for this)

	Can have multiple functions like:
	UpdateObjectThroughIndex
	UpdateObjectThroughCondition
	- name == "Ervin Howell" => username = "EH"

	since they would have different signatures and implementations I feel
*/

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func GetAllObjects(pathToFile string, obj interface{}) error {

	absPathToFile, _ := filepath.Abs(pathToFile)
	pathFile, err := os.Open(absPathToFile)
	if err != nil {
		return err
	}

	defer pathFile.Close()

	byteValue, err := io.ReadAll(pathFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, obj)
	if err != nil {
		return err
	}

	return nil
}

// When the JSON passed in is an array, then returning only the object
func GetObjectFromIndex[T any](pathToFile string, index int) (T, error) {

	var zeroVal T // This is the zero value for the datatype

	absPathToFile, _ := filepath.Abs(pathToFile)
	pathFile, err := os.Open(absPathToFile)
	if err != nil {
		return zeroVal, err
	}

	defer pathFile.Close()

	byteValue, err := io.ReadAll(pathFile)
	if err != nil {
		return zeroVal, err
	}

	var newObj []T

	err = json.Unmarshal(byteValue, &newObj)
	if err != nil {
		return zeroVal, err
	}

	if index < 0 && len(newObj) <= index {
		return zeroVal, fmt.Errorf("index %d not present in array", index)
	}

	return newObj[index], nil
}

func UpdateObjectThroughIndex(pathToFile string, obj interface{}) error {

	return nil
}

// Will see if I can work on updating modularly if possible
func SetAllObjects(pathToFile string, obj interface{}) error {
	absPathToFile, _ := filepath.Abs(pathToFile)

	pathFile, err := os.OpenFile(absPathToFile, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer pathFile.Close()

	updatedContents, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	_, err = pathFile.Write(updatedContents)
	if err != nil {
		return err
	}

	return nil
}
