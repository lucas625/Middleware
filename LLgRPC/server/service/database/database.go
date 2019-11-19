package database

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/lucas625/Middleware/LLgRPC/common/service/person"
	"github.com/lucas625/Middleware/utils"
)

// Database is a structure for holding the database.
//
// Members:
//  Persons - list of persons.
//
type Database struct {
	Persons *person.PersonList
	sync.Mutex
}

// AddPerson is a function for adding a person.
//
// Parameters:
//  p - the new person.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) AddPerson(p person.Person) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	return plist.AddPerson(p)
}

// RemovePerson is a function for removing a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) RemovePerson(id int) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	return plist.RemovePerson(id)
}

// GetPerson is a function for getting a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the person.
//
func (db *Database) GetPerson(id int) person.Person {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	if idx == -1 {
		fmt.Println("Person not found.")
		return *person.InitPerson("", 0, "", -1)
	}
	return plist.Persons[idx]
}

// SetPerson is a function for setting a person (keeping id).
//
// Parameters:
//  id - the id of the person.
//  p  - the new person.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) SetPerson(id int, p person.Person) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	if idx == -1 {
		fmt.Println("Person not found.")
		return false
	}
	p.SetID(id)
	plist.Persons[idx] = p
	return true
}

// SetName is a function for setting a person name.
//
// Parameters:
//  id    - the id of the person.
//  name  - the new name.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) SetName(id int, name string) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	if idx == -1 {
		fmt.Println("Person not found.")
		return false
	}
	plist.Persons[idx].SetName(name)
	return true
}

// SetAge is a function for setting a person age.
//
// Parameters:
//  id    - the id of the person.
//  age   - the new age.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) SetAge(id int, age int) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	if idx == -1 {
		fmt.Println("Person not found.")
		return false
	}
	plist.Persons[idx].SetAge(age)
	return true
}

// SetGender is a function for setting a person gender.
//
// Parameters:
//  id     - the id of the person.
//  gender - the new gender.
//
// Returns:
//  none
//
func (db *Database) SetGender(id int, gender string) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	if idx == -1 {
		fmt.Println("Person not found.")
		return false
	}
	plist.Persons[idx].SetGender(gender)
	return true
}

// DBToJSON is a function for converting the database to json.
//
// Parameters:
//  none
//
// Returns:
//  the database as byte.
//
func (db *Database) DBToJSON() []byte {
	db.Lock()
	defer db.Unlock()

	persons := make(map[string]interface{})
	persons["PersonList"] = make([]map[string]interface{}, len(db.Persons.Persons))
	for i := range db.Persons.Persons {
		persons["PersonList"].([]map[string]interface{})[i] = person.PersonToInterface(db.Persons.Persons[i])
	}
	persons["NextID"] = db.Persons.NextID
	file, err := json.MarshalIndent(persons, "", "	")
	utils.PrintError(err, "Unable to convert database to json.")
	return file
}

// LoadDatabase is a function for loading the database.
//
// Parameters:
//  plist - the new list of persons.
//
// Returns:
//  none
//
func (db *Database) LoadDatabase(plist *person.PersonList) {
	db.Lock()
	defer db.Unlock()
	db.Persons = plist
}

// InitDatabase is a function for initializing a database.
//
// Parameters:
//  none
//
// Returns:
//  the database.
//
func InitDatabase() *Database {
	pList := person.PersonList{Persons: make([]person.Person, 0, 100), NextID: 0}
	db := Database{Persons: &pList}
	return &db
}
