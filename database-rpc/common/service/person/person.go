package person

// PersonList is a structure for holding multiple persons.
//
// Members:
//  Persons - a list of persons.
//  NextID  - ID generator.
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
	ID := plist.NextID
	if ID >= cap(plist.Persons) {
		plist.Persons = plist.Persons[:cap(plist.Persons)]
	}
	p.SetID(ID)
	plist.NextID++
	plist.Persons = append(plist.Persons, p)
	return true
}

// RemovePerson is a function for adding a person.
//
// Parameters:
//  ID - the ID of the person.
//
// Returns:
//  a flag if went ok.
//
func (plist *PersonList) RemovePerson(ID int) bool {
	IDx := plist.GetPerson(ID)
	if IDx == -1 {
		return false
	}
	plist.Persons[IDx] = plist.Persons[len(plist.Persons)]
	plist.Persons = plist.Persons[:len(plist.Persons)-1]
	return true
}

// GetPerson is a function for getting a person IDx by ID.
//
// Parameters:
//  ID - the ID of the person.
//
// Returns:
//  IDx - the IDx of the person.
//
func (plist *PersonList) GetPerson(ID int) int {
	IDx := -1
	for i := range plist.Persons {
		if plist.Persons[i].GetID() == ID {
			IDx = i
		}
	}
	return IDx
}

// Person is a structure for saving a person's data.
//
// Members:
//  Name   - the Name of the person.
//  Age    - the Age of the person.
//  Gender - the Gender of the person.
//
type Person struct {
	Name   string
	Age    int
	Gender string
	ID     int
}

// PersonToInterface is a function for converting a person to interface.
//
// Parameters:
//  p - the person.
//
// Returns:
//  Name - the Name of the person.
//
func PersonToInterface(p Person) map[string]interface{} {
	per := make(map[string]interface{})
	per["Name"] = p.GetName()
	per["Age"] = p.GetAge()
	per["Gender"] = p.GetGender()
	per["ID"] = p.GetID()
	return per
}

// GetName is a function for getting a person Name.
//
// Parameters:
//  none
//
// Returns:
//  Name - the Name of the person.
//
func (p *Person) GetName() string {
	return p.Name
}

// GetAge is a function for getting a person Age.
//
// Parameters:
//  none
//
// Returns:
//  Age - the Age of the person.
//
func (p *Person) GetAge() int {
	return p.Age
}

// GetGender is a function for getting a person Gender.
//
// Parameters:
//  none
//
// Returns:
//  Gender - the Gender of the person.
//
func (p *Person) GetGender() string {
	return p.Gender
}

// GetID is a function for getting a person ID.
//
// Parameters:
//  none
//
// Returns:
//  ID - the ID of the person.
//
func (p *Person) GetID() int {
	return p.ID
}

// SetName is a function for setting a person Name.
//
// Parameters:
//  Name - the new Name.
//
// Returns:
//  none
//
func (p *Person) SetName(Name string) {
	p.Name = Name
}

// SetAge is a function for setting a person Age.
//
// Parameters:
//  Age - the Age of the person.
//
// Returns:
//  none
//
func (p *Person) SetAge(Age int) {
	p.Age = Age
}

// SetGender is a function for setting a person Gender.
//
// Parameters:
//  Gender - the Gender of the person.
//
// Returns:
//  none
//
func (p *Person) SetGender(Gender string) {
	p.Gender = Gender
}

// SetID is a function for setting a person ID.
//
// Parameters:
//  ID - the ID of the person.
//
// Returns:
//  none
//
func (p *Person) SetID(ID int) {
	p.ID = ID
}

// InitPerson is a function for initializing a person.
//
// Parameters:
//  Name   - the Name of the person.
//  Age    - the Age of the person.
//  Gender - the Gender of the person.
//  ID     - the ID of the person.
//
// Returns:
//  a person
//
func InitPerson(Name string, Age int, Gender string, ID int) *Person {
	p := Person{Name: Name, Age: Age, Gender: Gender, ID: ID}
	return &p
}
