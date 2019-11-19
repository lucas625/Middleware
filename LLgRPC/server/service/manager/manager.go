package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/lucas625/Middleware/LLgRPC/common/service/person"
	"github.com/lucas625/Middleware/LLgRPC/common/utils"
	"github.com/lucas625/Middleware/LLgRPC/server/service/database"
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
//  p - the new person.
//
// Returns:
//  a flag if went ok.
//
func (man *Manager) AddPerson(p person.Person) bool {
	return man.DB.AddPerson(p)
}

// RemovePerson is a function for removing a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  a flag if went ok.
//
func (man *Manager) RemovePerson(id int) bool {
	return man.DB.RemovePerson(id)
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
//  a flag if went ok.
//
func (man *Manager) SetPerson(id int, p person.Person) bool {
	return man.DB.SetPerson(id, p)
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
//  a flag if went ok.
//
func (man *Manager) SetName(id int, name string) bool {
	return man.DB.SetName(id, name)
}

// SetAge is a function for setting a person is age.
//
// Parameters:
//  id  - the id of the person.
//  age - the age of the person.
//
// Returns:
//  a flag if went ok.
//
func (man *Manager) SetAge(id int, age int) bool {
	return man.DB.SetAge(id, age)
}

// SetGender is a function for setting a person is gender.
//
// Parameters:
//  id     - the id of the person.
//  gender - the gender of the person.
//
// Returns:
//  a flag if went ok.
//
func (man *Manager) SetGender(id int, gender string) bool {
	return man.DB.SetGender(id, gender)
}

// List is a function to list all data of the database.
//
// Parameters:
//  none
//
// Returns:
//  the list of persons on the database.
//
func (man *Manager) List() map[string]interface{} {
	var interfaceC map[string]interface{}
	err := json.Unmarshal(man.DB.DBToJSON(), &interfaceC)
	if err != nil {
		fmt.Println("Failed to list", err)
		return make(map[string]interface{})
	}
	return interfaceC
}

// Write is a function to write all data of the database.
//
// Parameters:
//  outpath - path to write the database.
//
// Returns:
//  a flag if went ok.
//
func (man *Manager) Write(outPath string) bool {
	fmt.Println("Writing database")
	// creating the json
	file := man.DB.DBToJSON()
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, "database.json"))
	if err != nil {
		fmt.Println("Failed to find path")
		return false
	}
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		if err != nil {
			fmt.Println("Failed to make dir on writer.")
			return false
		}
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	if err != nil {
		fmt.Println("Failed to write.")
		return false
	}
	return true
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
	if utils.PathExists(inPath) {
		// opening the file
		dbFile, err := os.Open(inPath)
		utils.PrintError(err, "Unable to open database.")
		// converting to PersonList
		byteDatabase, err := ioutil.ReadAll(dbFile)
		utils.PrintError(err, "Unable to convert database file to bytes.")

		var pListMap map[string]interface{}

		err = json.Unmarshal(byteDatabase, &pListMap)

		pListIF := pListMap["PersonList"].([]interface{})
		nextID := int(pListMap["NextID"].(float64))

		personlist := make([]person.Person, len(pListIF))

		for i := range pListIF {
			paux := person.InitPerson(pListIF[i].(map[string]interface{})["name"].(string),
				int(pListIF[i].(map[string]interface{})["age"].(float64)),
				pListIF[i].(map[string]interface{})["gender"].(string),
				int(pListIF[i].(map[string]interface{})["id"].(float64)))
			personlist[i] = *paux
		}

		pList := person.PersonList{Persons: personlist, NextID: nextID}

		utils.PrintError(err, "Failed to unmarshal database.")
		// loading the database
		man.DB.LoadDatabase(&pList)
		fmt.Println("Database loaded.")
	} else {
		man.DB = database.InitDatabase()
		fmt.Println("Unable to find database.json.")
	}
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
