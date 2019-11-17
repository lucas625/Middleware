package manager

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/lucas625/Middleware/LLgRPC/common/utils"
	"github.com/lucas625/Middleware/LLgRPC/server/service/database"
	"github.com/lucas625/Middleware/LLgRPC/server/service/person"
)

// Manager is a structure for managing the database.
//
// Members:
//  none
//
type Manager struct {
	DB *database.Database
}

// AddPerson is a function for adding a person.
//
// Parameters:
//  p  - the new person.
//
// Returns:
//  none
//
func (man *Manager) AddPerson(p person.Person) {
	man.DB.AddPerson(p)
}

// RemovePerson is a function for removing a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  none
//
func (man *Manager) RemovePerson(id int) {
	man.DB.RemovePerson(id)
}

// GetPerson is a function for getting a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the person.
//
func (man *Manager) GetPerson(id int) person.Person {
	return man.DB.GetPerson(id)
}

// SetPerson is a function for setting a person.
//
// Parameters:
//  id - the id of the person.
//  p  - the new person.
//
// Returns:
//  none
//
func (man *Manager) SetPerson(id int, p person.Person) {
	man.DB.SetPerson(id, p)
}

// GetName is a function for getting a person is name.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the name of the person.
//
func (man *Manager) GetName(id int) string {
	person := man.DB.GetPerson(id)
	return person.GetName()
}

// GetAge is a function for getting a person is age.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the age of the person.
//
func (man *Manager) GetAge(id int) int {
	person := man.DB.GetPerson(id)
	return person.GetAge()
}

// GetGender is a function for getting a person is gender.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the gender of the person.
//
func (man *Manager) GetGender(id int) string {
	person := man.DB.GetPerson(id)
	return person.GetGender()
}

// SetName is a function for setting a person is name.
//
// Parameters:
//  id   - the id of the person.
//  name - the name of the person.
//
// Returns:
//  none
//
func (man *Manager) SetName(id int, name string) {
	man.DB.SetName(id, name)
}

// SetAge is a function for setting a person is age.
//
// Parameters:
//  id  - the id of the person.
//  age - the age of the person.
//
// Returns:
//  none
//
func (man *Manager) SetAge(id int, age int) {
	man.DB.SetAge(id, age)
}

// SetGender is a function for setting a person is gender.
//
// Parameters:
//  id     - the id of the person.
//  gender - the gender of the person.
//
// Returns:
//  none
//
func (man *Manager) SetGender(id int, gender string) {
	man.DB.SetGender(id, gender)
}

// Write is a function to write all data of the database.
//
// Parameters:
//  none
//
// Returns:
//  none
//
func (man *Manager) Write(outPath string) {
	// creating the json
	file := man.DB.DBToJson()
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, "database.json"))
	utils.PrintError(err, "Unable to get database is absolute path.")
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		utils.PrintError(err, "Unable to create dirs.")
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	utils.PrintError(err, "Unable to write database.")
}

// Load is a function to load all database as json.
//
// Parameters:
//  inPath - path to the input file.
//
// Returns:
//  none
//
func (man *Manager) Load(inPath string) {
	// opening the file
	dbFile, err := os.Open(inPath)
	utils.PrintError(err, "Unable to open database.")
	// converting to PersonList
	byteDatabase, err := ioutil.ReadAll(dbFile)
	utils.PrintError(err, "Unable to convert database file to bytes.")
	var pList person.PersonList
	err = json.Unmarshal(byteDatabase, &pList)
	utils.PrintError(err, "Failed to unmarshal database.")
	// loading the database
	man.DB.LoadDatabase(&pList)
}

// InitManager is a function for initializing the manager.
//
// Parameters:
//  db - the database.
//
// Returns:
//  the manager
//
func InitManager(db *database.Database) *Manager {
	man := Manager{DB: db}
	return &man
}
