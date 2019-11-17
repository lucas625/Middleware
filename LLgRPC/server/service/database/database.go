package database

import (
	"encoding/json"
	"sync"

	"github.com/lucas625/Middleware/LLgRPC/server/service/person"
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
//  p  - the new person.
//
// Returns:
//  none
//
func (db *Database) AddPerson(p person.Person) {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	plist.AddPerson(p)
}

// RemovePerson is a function for removing a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  none
//
func (db *Database) RemovePerson(id int) {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	plist.RemovePerson(id)
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
	return plist.Persons[idx]
}

// SetPerson is a function for setting a person (keeping id).
//
// Parameters:
//  id - the id of the person.
//  p  - the new person.
//
// Returns:
//  none
//
func (db *Database) SetPerson(id int, p person.Person) {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	p.SetID(id)
	plist.Persons[idx] = p
}

// SetName is a function for setting a person name.
//
// Parameters:
//  id    - the id of the person.
//  name  - the new name.
//
// Returns:
//  none
//
func (db *Database) SetName(id int, name string) {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	plist.Persons[idx].SetName(name)
}

// SetAge is a function for setting a person age.
//
// Parameters:
//  id    - the id of the person.
//  age   - the new age.
//
// Returns:
//  none
//
func (db *Database) SetAge(id int, age int) {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	plist.Persons[idx].SetAge(age)
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
func (db *Database) SetGender(id int, gender string) {
	db.Lock()
	defer db.Unlock()
	plist := db.Persons
	idx := plist.GetPerson(id)
	plist.Persons[idx].SetGender(gender)
}

// DBToJson is a function for converting the database to json.
//
// Parameters:
//  none
//
// Returns:
//  the database as byte.
//
func (db *Database) DBToJson() []byte {
	db.Lock()
	defer db.Unlock()

	pl := db.Persons
	file, err := json.MarshalIndent(*pl, "", "	")
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
	pList := person.PersonList{Persons: make([]person.Person, 100), NextID: 0}
	db := Database{Persons: &pList}
	return &db
}
