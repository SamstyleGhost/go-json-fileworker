# About
This library makes it easier to work with local JSON files in Go. Has multiple handy abstractions to get, set & update JSON objects and arrays.  
It is currently a Work In Progress, and I built it as I did not really find any library of the same exact working. If you just want to work with JSON objects in code, checkout:
- [gjson](https://github.com/tidwall/gjson)
- [sjson](https://github.com/tidwall/sjson)

## Features currently built
  - [GetAllObjects](#getallobjects)
  - [GetObjectFromIndex](#getobjectfromindex)
  - [SetAllObjects](#setallobjects)
  - [AppendObjectToArray](#appendobjecttoarray)
  - [AppendObjectToArrayDirect](#appendobjecttoarraydirect)

#### Path Note: In all functions, the ***filePath*** works with both absolute as well as relative paths 

### GetAllObjects
- Returns all the contents from the file
- This function unmarshals the object directly into the variable passed (Does require passing the var by reference.)
- If datatype is unknown, passing in an interface{} would still give results.
```go
func main() {
  var users []models.Users // Replace with the datatype required.

  err := GetAllObjects(filePath, &users)
  // Handle errors

  b, _ := json.MarshalIndent(users, "", "  ")
  fmt.Print(string(b))
}
```

### GetObjectFromIndex
- Returns a single object from an array of objects
- Uses generics to take care of typesafety & autocomplete
- Would throw an error if the index passed does not exist
- If datatype is unknown, can pass in *any* or discard it altogether
```go
func main() {

  v, err := GetObjectFromIndex[models.Users](filePath, 5)
  // Handle errors

  fmt.Print(v.Name)
}
```

### SetAllObjects
- Clears out the entire contents of the file and replaces it with the given contents
```go
// To be added
```

### AppendObjectToArray
- Appends an object to the end of the array.
- Follows the RMW (Read-Modify-Write) approach. Reads the entire file into memory, modifies it and then writes the file contents back.
- <span style="color: #CC0000">Currently only works with the entire object array. Working on appending objects to array fields.<span>
```go
func main() {
  
  obj := models.Users{
    ID:       1,
    Name:     "Leanne Graham",
    Username: "Bret",
    Email:    "Sincere@april.biz",
    Address: models.Address{
      // ...
    }
  }

  err = AppendObjectToArray(filePath, obj)

  // optional generic can also be added if you are sure that the other objects in the array are also of the same type
  // err = AppendObjectToArray[models.Users](filePath, obj)

  // Handle errors
}
```

### AppendObjectToArrayDirect
- Does the same thing as [AppendObjectToArray](#appendobjecttoarray)
- But instead of reading all the filecontents into memory, appends it directly using some manipulation.
- Is recommended when the filesize is too large to be read into memory.
```go
func main() {
  
  obj := models.Users{
    ID:       1,
    Name:     "Leanne Graham",
    Username: "Bret",
    Email:    "Sincere@april.biz",
    Address: models.Address{
      // ...
    }
  }

  err = AppendObjectToArrayDirect(filePath, obj)
  // Handle errors
}

```

---
Things to note:
- If need more control over updates, it is better to GetAllObjects, run the updates and then SetAllObjects
- The generics are used for type safety and JSON structure consistency. Can use *any* if the structure is unknown or doesnt matter


## TODO:
- Check wrongdatatypes for GetAllObjects
- Test and add code example for SetAllObjects 