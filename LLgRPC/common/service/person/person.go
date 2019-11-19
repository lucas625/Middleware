package person

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
//  a flag if went ok.
//
func (plist *PersonList) AddPerson(p Person) bool {
	id := plist.NextID
	if id >= cap(plist.Persons) {
		plist.Persons = plist.Persons[:cap(plist.Persons)]
	}
	p.SetID(id)
	plist.NextID++
	plist.Persons = append(plist.Persons, p)
	return true
}

// RemovePerson is a function for adding a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  a flag if went ok.
//
func (plist *PersonList) RemovePerson(id int) bool {
	idx := plist.GetPerson(id)
	if idx == -1 {
		return false
	}
	plist.Persons[idx] = plist.Persons[len(plist.Persons)]
	plist.Persons = plist.Persons[:len(plist.Persons)-1]
	return true
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

// PersonToInterface is a function for converting a person to interface.
//
// Parameters:
//  p - the person.
//
// Returns:
//  name - the name of the person.
//
func PersonToInterface(p Person) map[string]interface{} {
	per := make(map[string]interface{})
	per["name"] = p.GetName()
	per["age"] = p.GetAge()
	per["gender"] = p.GetGender()
	per["id"] = p.GetID()
	return per
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
