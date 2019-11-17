package person

import (
	"errors"

	"github.com/lucas625/Middleware/utils"
)

// PersonList is a structure for holding multiple persons.
//
// Members:
//  Persons - a list of persons.
//  NextID  - id generator.
//
type PersonList struct {
	Persons []Person
	NextID  int
}

// AddPerson is a function for adding a person.
//
// Parameters:
//  p - the person
//
// Returns:
//  none
//
func (plist *PersonList) AddPerson(p Person) {
	id := plist.NextID
	if id >= cap(plist.Persons) {
		plist.Persons = plist.Persons[:cap(plist.Persons)]
	}
	p.SetID(id)
	plist.NextID += 1
	plist.Persons = append(plist.Persons, p)
}

// RemovePerson is a function for adding a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  none
//
func (plist *PersonList) RemovePerson(id int) {
	idx := plist.GetPerson(id)
	plist.Persons[idx] = plist.Persons[len(plist.Persons)]
	plist.Persons = plist.Persons[:len(plist.Persons)-1]
}

// GetPerson is a function for getting a person idx by id.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  idx - the idx of the person.
//
func (plist *PersonList) GetPerson(id int) int {
	idx := -1
	for i := range plist.Persons {
		if plist.Persons[i].GetID() == id {
			idx = i
		}
	}
	if idx == -1 {
		utils.PrintError(errors.New("Unable to find person."), "Invalid id for person.")
	}
	return idx
}

// Person is a structure for saving a person's data.
//
// Members:
//  name   - the name of the person.
//  age    - the age of the person.
//  gender - the gender of the person.
//
type Person struct {
	name   string
	age    int
	gender string
	id     int
}

// GetName is a function for getting a person name.
//
// Parameters:
//  none
//
// Returns:
//  name - the name of the person.
//
func (p *Person) GetName() string {
	return p.name
}

// GetAge is a function for getting a person age.
//
// Parameters:
//  none
//
// Returns:
//  age - the age of the person.
//
func (p *Person) GetAge() int {
	return p.age
}

// GetGender is a function for getting a person gender.
//
// Parameters:
//  none
//
// Returns:
//  gender - the gender of the person.
//
func (p *Person) GetGender() string {
	return p.gender
}

// GetID is a function for getting a person id.
//
// Parameters:
//  none
//
// Returns:
//  id - the id of the person.
//
func (p *Person) GetID() int {
	return p.id
}

// SetName is a function for setting a person name.
//
// Parameters:
//  name - the new name.
//
// Returns:
//  none
//
func (p *Person) SetName(name string) {
	p.name = name
}

// SetAge is a function for setting a person age.
//
// Parameters:
//  age - the age of the person.
//
// Returns:
//  none
//
func (p *Person) SetAge(age int) {
	p.age = age
}

// SetGender is a function for setting a person gender.
//
// Parameters:
//  gender - the gender of the person.
//
// Returns:
//  none
//
func (p *Person) SetGender(gender string) {
	p.gender = gender
}

// SetID is a function for setting a person id.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  none
//
func (p *Person) SetID(id int) {
	p.id = id
}

// InitPerson is a function for initializing a person.
//
// Parameters:
//  name   - the name of the person.
//  age    - the age of the person.
//  gender - the gender of the person.
//  id     - the id of the person.
//
// Returns:
//  a person
//
func InitPerson(name string, age int, gender string, id int) *Person {
	p := Person{name: name, age: age, gender: gender, id: id}
	return &p
}
