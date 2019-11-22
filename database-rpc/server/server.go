package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	//"github.com/lucas625/Middleware/utils"
	//"github.com/lucas625/Middleware/mom-rpc/rpc/multiplicador/impl"
	//"github.com/lucas625/Middleware/LLgRPC/server/service/database"
	//"github.com/lucas625/Middleware/LLgRPC/server/service/manager"
	//"github.com/lucas625/Middleware/database-rpc/service/manager"
	//"github.com/lucas625/Middleware/database-rpc/service/database"
	//"github.com/lucas625/Middleware/database-rpc/common/service/person"


	"encoding/json"
	"sync"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	//"github.com/lucas625/Middleware/database-rpc/common/service/person"
	"github.com/lucas625/Middleware/LLgRPC/common/utils"
	//"github.com/lucas625/Middleware/database-rpc/service/database"
)

func main() {

	// create new instance of database
	manager1 := new(manager)

	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Database", manager1)

	// associate a http handler to servidor
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, err := net.Listen("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "Servidor não inicializado")

	// wait for calls
	fmt.Println("Servidor pronto (RPC-HTTP) ...\n")
	http.Serve(l, nil)
}

// manager is a structure for managing the 
//
// Members:
//  none
//
type manager struct {
	DB *Database
}

// Addperson is a function for adding a 
//
// Parameters:
//  p - the new 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) Addperson(p person) bool {
	return man.DB.Addperson(p)
}

// Removeperson is a function for removing a 
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) Removeperson(id int) bool {
	return man.DB.Removeperson(id)
}

// Getperson is a function for getting a 
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  the 
//
func (man *manager) Getperson(id int) person {
	return man.DB.Getperson(id)
}

// Setperson is a function for setting a 
//
// Parameters:
//  id - the id of the 
//  p  - the new 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) Setperson(id int, p person) bool {
	return man.DB.Setperson(id, p)
}

// GetName is a function for getting a person is name.
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  the name of the 
//
func (man *manager) GetName(id int) string {
	person := man.DB.Getperson(id)
	return person.GetName()
}

// GetAge is a function for getting a person is age.
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  the age of the 
//
func (man *manager) GetAge(id int) int {
	person := man.DB.Getperson(id)
	return person.GetAge()
}

// GetGender is a function for getting a person is gender.
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  the gender of the 
//
func (man *manager) GetGender(id int) string {
	person := man.DB.Getperson(id)
	return person.GetGender()
}

// SetName is a function for setting a person is name.
//
// Parameters:
//  id   - the id of the 
//  name - the name of the 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) SetName(id int, name string) bool {
	return man.DB.SetName(id, name)
}

// SetAge is a function for setting a person is age.
//
// Parameters:
//  id  - the id of the 
//  age - the age of the 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) SetAge(id int, age int) bool {
	return man.DB.SetAge(id, age)
}

// SetGender is a function for setting a person is gender.
//
// Parameters:
//  id     - the id of the 
//  gender - the gender of the 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) SetGender(id int, gender string) bool {
	return man.DB.SetGender(id, gender)
}

// List is a function to list all data of the 
//
// Parameters:
//  none
//
// Returns:
//  the list of persons on the 
//
func (man *manager) List() map[string]interface{} {
	var interfaceC map[string]interface{}
	err := json.Unmarshal(man.DB.DBToJSON(), &interfaceC)
	if err != nil {
		fmt.Println("Failed to list", err)
		return make(map[string]interface{})
	}
	return interfaceC
}

// Write is a function to write all data of the 
//
// Parameters:
//  outpath - path to write the 
//
// Returns:
//  a flag if went ok.
//
func (man *manager) Write(outPath string) bool {
	fmt.Println("Writing database")
	// creating the json
	file := man.DB.DBToJSON()
	// getting the right path
	filePath, err := filepath.Abs(path.Join(outPath, "json"))
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
func (man *manager) Load(inPath string) {
	if utils.PathExists(inPath) {
		// opening the file
		dbFile, err := os.Open(inPath)
		utils.PrintError(err, "Unable to open ")
		// converting to personList
		byteDatabase, err := ioutil.ReadAll(dbFile)
		utils.PrintError(err, "Unable to convert database file to bytes.")

		var pListMap map[string]interface{}

		err = json.Unmarshal(byteDatabase, &pListMap)

		pListIF := pListMap["personList"].([]interface{})
		nextID := int(pListMap["NextID"].(float64))

		personlist := make([]person, len(pListIF))

		for i := range pListIF {
			paux := Initperson(pListIF[i].(map[string]interface{})["name"].(string),
				int(pListIF[i].(map[string]interface{})["age"].(float64)),
				pListIF[i].(map[string]interface{})["gender"].(string),
				int(pListIF[i].(map[string]interface{})["id"].(float64)))
			personlist[i] = *paux
		}

		pList := personList{persons: personlist, NextID: nextID}

		utils.PrintError(err, "Failed to unmarshal ")
		// loading the database
		man.DB.LoadDatabase(&pList)
		fmt.Println("Database loaded.")
	} else {
		man.DB = InitDatabase()
		fmt.Println("Unable to find json.")
	}
}

// Initmanager is a function for initializing the manager.
//
// Parameters:
//  db - the 
//
// Returns:
//  the manager
//
func Initmanager(db *Database) *manager {
	man := manager{DB: db}
	return &man
}


type personList struct {
	persons []person
	NextID  int
}

// Addperson is a function for adding a 
//
// Parameters:
//  p - the person
//
// Returns:
//  a flag if went ok.
//
func (plist *personList) Addperson(p person) bool {
	id := plist.NextID
	if id >= cap(plist.persons) {
		plist.persons = plist.persons[:cap(plist.persons)]
	}
	p.SetID(id)
	plist.NextID++
	plist.persons = append(plist.persons, p)
	return true
}

// Removeperson is a function for adding a 
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  a flag if went ok.
//
func (plist *personList) Removeperson(id int) bool {
	idx := plist.Getperson(id)
	if idx == -1 {
		return false
	}
	plist.persons[idx] = plist.persons[len(plist.persons)]
	plist.persons = plist.persons[:len(plist.persons)-1]
	return true
}

// Getperson is a function for getting a person idx by id.
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  idx - the idx of the 
//
func (plist *personList) Getperson(id int) int {
	idx := -1
	for i := range plist.persons {
		if plist.persons[i].GetID() == id {
			idx = i
		}
	}
	return idx
}

// person is a structure for saving a person's data.
//
// Members:
//  name   - the name of the 
//  age    - the age of the 
//  gender - the gender of the 
//
type person struct {
	name   string
	age    int
	gender string
	id     int
}

// personToInterface is a function for converting a person to interface.
//
// Parameters:
//  p - the 
//
// Returns:
//  name - the name of the 
//
func personToInterface(p person) map[string]interface{} {
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
//  name - the name of the 
//
func (p *person) GetName() string {
	return p.name
}

// GetAge is a function for getting a person age.
//
// Parameters:
//  none
//
// Returns:
//  age - the age of the 
//
func (p *person) GetAge() int {
	return p.age
}

// GetGender is a function for getting a person gender.
//
// Parameters:
//  none
//
// Returns:
//  gender - the gender of the 
//
func (p *person) GetGender() string {
	return p.gender
}

// GetID is a function for getting a person id.
//
// Parameters:
//  none
//
// Returns:
//  id - the id of the 
//
func (p *person) GetID() int {
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
func (p *person) SetName(name string) {
	p.name = name
}

// SetAge is a function for setting a person age.
//
// Parameters:
//  age - the age of the 
//
// Returns:
//  none
//
func (p *person) SetAge(age int) {
	p.age = age
}

// SetGender is a function for setting a person gender.
//
// Parameters:
//  gender - the gender of the 
//
// Returns:
//  none
//
func (p *person) SetGender(gender string) {
	p.gender = gender
}

// SetID is a function for setting a person id.
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  none
//
func (p *person) SetID(id int) {
	p.id = id
}

// Initperson is a function for initializing a 
//
// Parameters:
//  name   - the name of the 
//  age    - the age of the 
//  gender - the gender of the 
//  id     - the id of the 
//
// Returns:
//  a person
//
func Initperson(name string, age int, gender string, id int) *person {
	p := person{name: name, age: age, gender: gender, id: id}
	return &p
}

type Database struct {
	persons *personList
	sync.Mutex
}

// Addperson is a function for adding a 
//
// Parameters:
//  p - the new 
//
// Returns:
//  a flag if went ok.
//
func (db *Database) Addperson(p person) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	return plist.Addperson(p)
}

// Removeperson is a function for removing a 
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  a flag if went ok.
//
func (db *Database) Removeperson(id int) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	return plist.Removeperson(id)
}

// Getperson is a function for getting a 
//
// Parameters:
//  id - the id of the 
//
// Returns:
//  the 
//
func (db *Database) Getperson(id int) person {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	idx := plist.Getperson(id)
	if idx == -1 {
		fmt.Println("person not found.")
		return *Initperson("", 0, "", -1)
	}
	return plist.persons[idx]
}

// Setperson is a function for setting a person (keeping id).
//
// Parameters:
//  id - the id of the 
//  p  - the new 
//
// Returns:
//  a flag if went ok.
//
func (db *Database) Setperson(id int, p person) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	idx := plist.Getperson(id)
	if idx == -1 {
		fmt.Println("person not found.")
		return false
	}
	p.SetID(id)
	plist.persons[idx] = p
	return true
}

// SetName is a function for setting a person name.
//
// Parameters:
//  id    - the id of the 
//  name  - the new name.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) SetName(id int, name string) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	idx := plist.Getperson(id)
	if idx == -1 {
		fmt.Println("person not found.")
		return false
	}
	plist.persons[idx].SetName(name)
	return true
}

// SetAge is a function for setting a person age.
//
// Parameters:
//  id    - the id of the 
//  age   - the new age.
//
// Returns:
//  a flag if went ok.
//
func (db *Database) SetAge(id int, age int) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	idx := plist.Getperson(id)
	if idx == -1 {
		fmt.Println("person not found.")
		return false
	}
	plist.persons[idx].SetAge(age)
	return true
}

// SetGender is a function for setting a person gender.
//
// Parameters:
//  id     - the id of the 
//  gender - the new gender.
//
// Returns:
//  none
//
func (db *Database) SetGender(id int, gender string) bool {
	db.Lock()
	defer db.Unlock()
	plist := db.persons
	idx := plist.Getperson(id)
	if idx == -1 {
		fmt.Println("person not found.")
		return false
	}
	plist.persons[idx].SetGender(gender)
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
	persons["personList"] = make([]map[string]interface{}, len(db.persons.persons))
	for i := range db.persons.persons {
		persons["personList"].([]map[string]interface{})[i] = personToInterface(db.persons.persons[i])
	}
	persons["NextID"] = db.persons.NextID
	file, err := json.MarshalIndent(persons, "", "	")
	utils.PrintError(err, "Unable to convert database to json.")
	return file
}

// LoadDatabase is a function for loading the 
//
// Parameters:
//  plist - the new list of persons.
//
// Returns:
//  none
//
func (db *Database) LoadDatabase(plist *personList) {
	db.Lock()
	defer db.Unlock()
	db.persons = plist
}

// InitDatabase is a function for initializing a 
//
// Parameters:
//  none
//
// Returns:
//  the 
//
func InitDatabase() *Database {
	pList := personList{persons: make([]person, 0, 100), NextID: 0}
	db := Database{persons: &pList}
	return &db
}