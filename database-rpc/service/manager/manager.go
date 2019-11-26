package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/lucas625/Middleware/LLgRPC/common/utils"
	"github.com/lucas625/Middleware/database-rpc/common/service/person"
	"github.com/lucas625/Middleware/database-rpc/service/database"
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
//  p     - the new person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) AddPerson(p person.Person, reply *bool) error {
	*reply = man.DB.AddPerson(p)
	return nil
}

// RemovePerson is a function for removing a person.
//
// Parameters:
//  id    - the id of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) RemovePerson(id int, reply *bool) error {
	*reply = man.DB.RemovePerson(id)
	return nil
}

// GetPerson is a function for getting a person.
//
// Parameters:
//  id    - the id of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) GetPerson(id int, reply *person.Person) error {
	*reply = man.DB.GetPerson(id)
	return nil
}

// SetPerson is a function for setting a person.
//
// Parameters:
//  id    - the id of the person.
//  p     - the new person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) SetPerson(req map[string]interface{}, reply *bool) error {
	id := int(req["id"].(float64))
	p := req["age"].(person.Person)
	*reply = man.DB.SetPerson(id, p)
	return nil
}

// GetName is a function for getting a person is name.
//
// Parameters:
//  id    - the id of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) GetName(id int, reply *string) error {
	person := man.DB.GetPerson(id)
	*reply = person.GetName()
	return nil
}

// GetAge is a function for getting a person is age.
//
// Parameters:
//  id    - the id of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) GetAge(id int, reply *int) error {
	person := man.DB.GetPerson(id)
	*reply = person.GetAge()
	return nil
}

// GetGender is a function for getting a person is gender.
//
// Parameters:
//  id    - the id of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) GetGender(id int, reply *string) error {
	person := man.DB.GetPerson(id)
	*reply = person.GetGender()
	return nil
}

// SetName is a function for setting a person is name.
//
// Parameters:
//  id    - the id of the person.
//  name  - the name of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) SetName(req map[string]interface{}, reply *bool) error {
	id := int(req["id"].(float64))
	name := req["name"].(string)
	*reply = man.DB.SetName(id, name)
	return nil
}

// SetAge is a function for setting a person is age.
//
// Parameters:
//  id    - the id of the person.
//  age   - the age of the person.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) SetAge(req map[string]interface{}, reply *bool) error {
	id := int(req["id"].(float64))
	age := int(req["age"].(float64))
	*reply = man.DB.SetAge(id, age)
	return nil
}

// SetGender is a function for setting a person is gender.
//
// Parameters:
//  id     - the id of the person.
//  gender - the gender of the person.
//  reply  - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) SetGender(req map[string]interface{}, reply *bool) error {
	id := int(req["id"].(float64))
	gender := req["gender"].(string)
	*reply = man.DB.SetGender(id, gender)
	return nil
}

// List is a function to list all data of the database.
//
// Parameters:
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) List(req interface{}, reply *map[string]interface{}) error {
	var interfaceC map[string]interface{}
	err := json.Unmarshal(man.DB.DBToJSON(), &interfaceC)
	if err != nil {
		fmt.Println("Failed to list", err)
		*reply = make(map[string]interface{})
	}
	*reply = interfaceC
	return nil
}

// Write is a function to write all data of the database.
//
// Parameters:
//  outpath - path to write the database.
//  reply - the return of the operation.
//
// Returns:
//  none
//
func (man *Manager) Write(outPath string, reply *bool) error {
	fmt.Println("Writing database")
	// creating the json
	file := man.DB.DBToJSON()
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, "database.json"))
	if err != nil {
		fmt.Println("Failed to find path")
		*reply = false
	}
	// creating the folder if it doesn't exists.
	if !utils.PathExists(filePath) {
		err = os.MkdirAll(outPath, 0700)
		if err != nil {
			fmt.Println("Failed to make dir on writer.")
			*reply = false
		}
	}
	// writing
	err = ioutil.WriteFile(filePath, file, 0700)
	if err != nil {
		fmt.Println("Failed to write.")
		*reply = false
	}
	*reply = true
	return nil
}

// Load is a function to load all database as json.
//
// Parameters:
//  inPath - path to the input file.
//
// Returns:
//  none
//
func (man *Manager) Load(inPath string, reply *bool) error {
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
	*reply = true
	return nil
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
