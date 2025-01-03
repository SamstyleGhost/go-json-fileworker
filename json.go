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

/*
	Additional info:
	If the user passes a certain type in the generics and the object returned does not have certain fields, then they would be replaced with their zero values, unless the json is specified to omitempty. If the user doesnt want to conform to a particular type,they are advised to use any as the the type
*/

var (
	// This is taken from the jsoniter library docs.
	// It has the configs for the jsoniter.ConfigCompatibleWithStandardLibrary API + DisallowUnknowFields set to true, so as to have typesafety in case the JSON is of different type than the type given
	json = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		DisallowUnknownFields:  true,
	}.Froze()
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

/*
Returns the individual object at the given index.
Takes in the type that should be returned (This would help in typesafety as well as autocomplete)
Returns error when:
1. The file is not found
2. There is some error reading the file contents
3. The actual JSON content is not an array (need it to be an array to have an index for this one)
4. The content received is not of the type passed in

* Can pass in any as the type in the function if type is not known
* The function works with optional field values as well
*/
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

// Appends entire object to the array in file
// Does not yet work on arrays in individual fields
func AppendObjectToArray[T any](pathToFile string, obj T) error {

	absPathToFile, _ := filepath.Abs(pathToFile)

	pathFile, err := os.OpenFile(absPathToFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer pathFile.Close()

	byteValue, err := io.ReadAll(pathFile)
	if err != nil {
		return err
	}

	var objects []T
	err = json.Unmarshal(byteValue, &objects)
	if err != nil {
		return err
	}

	objects = append(objects, obj)

	updatedContents, err := json.MarshalIndent(objects, "", "  ")
	if err != nil {
		return err
	}

	if err = pathFile.Truncate(0); err != nil {
		return err
	}
	if _, err := pathFile.Seek(0, 0); err != nil { // Move to the beginning of the file
		return err
	}

	_, err = pathFile.Write(updatedContents)
	if err != nil {
		return err
	}

	return nil

}
