package jsonfileworker

import (
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

	since they would have different signatures and implementations I feel
*/

func GetAllObjects(pathToFile string, v interface{}) error {

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

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, v)
	if err != nil {
		return err
	}

	return nil
}

// Will see if I can work on updating modularly if possible
func SetAllObjects(pathToFile string, v interface{}) error {
	absPathToFile, _ := filepath.Abs(pathToFile)

	pathFile, err := os.OpenFile(absPathToFile, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer pathFile.Close()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	updatedContents, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = pathFile.Write(updatedContents)
	if err != nil {
		return err
	}

	return nil
}
