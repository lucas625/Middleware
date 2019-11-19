package proxies

import (
	"fmt"

	"github.com/lucas625/Middleware/LLgRPC/client/distribution/requestor"
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/absoluteobjectreference"
	"github.com/lucas625/Middleware/LLgRPC/common/service/person"
	"github.com/lucas625/Middleware/LLgRPC/common/utils"
)

// ManagerProxy is a struct that holds the data need to contact the server
//
// Members:
//  AOR - a absolute object reference.
//
type ManagerProxy struct {
	AOR absoluteobjectreference.AOR
}

// Write is a function to write all data of the database.
//
// Parameters:
//  outpath - path to write the database.
//
// Returns:
//  none
//
func (proxy ManagerProxy) Write(path string) {
	param := make([]interface{}, 1)
	param[0] = path
	rq := utils.Request{Op: "Write", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	fmt.Println(reply)
}

// AddPerson is a function for adding a person.
//
// Parameters:
//  p - the new person.
//
// Returns:
//  a flag if went ok.
//
func (proxy ManagerProxy) AddPerson(p person.Person) bool {
	param := make([]interface{}, 1)
	param[0] = person.PersonToInterface(p)
	rq := utils.Request{Op: "AddPerson", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(bool)
}

// RemovePerson is a function for removing a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  a flag if went ok.
//
func (proxy ManagerProxy) RemovePerson(id int) bool {
	param := make([]interface{}, 1)
	param[0] = id
	rq := utils.Request{Op: "RemovePerson", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(bool)
}

// GetPerson is a function for getting a person.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the person.
//
func (proxy ManagerProxy) GetPerson(id int) person.Person {
	param := make([]interface{}, 1)
	param[0] = id
	rq := utils.Request{Op: "GetPerson", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	rp := reply[0].(map[string]interface{})
	p := person.InitPerson(
		rp["name"].(string),
		int(rp["age"].(float64)),
		rp["gender"].(string),
		int(rp["id"].(float64)))
	return *p
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
func (proxy ManagerProxy) SetPerson(id int, p person.Person) bool {
	param := make([]interface{}, 2)
	param[0] = id
	param[1] = person.PersonToInterface(p)
	rq := utils.Request{Op: "SetPerson", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(bool)
}

// GetName is a function for getting a person is name.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the name of the person.
//
func (proxy ManagerProxy) GetName(id int) string {
	param := make([]interface{}, 1)
	param[0] = id
	rq := utils.Request{Op: "GetName", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(string)
}

// GetAge is a function for getting a person is age.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the age of the person.
//
func (proxy ManagerProxy) GetAge(id int) int {
	param := make([]interface{}, 1)
	param[0] = id
	rq := utils.Request{Op: "GetAge", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return int(reply[0].(float64))
}

// GetGender is a function for getting a person is gender.
//
// Parameters:
//  id - the id of the person.
//
// Returns:
//  the gender of the person.
//
func (proxy ManagerProxy) GetGender(id int) string {
	param := make([]interface{}, 1)
	param[0] = id
	rq := utils.Request{Op: "GetGender", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(string)
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
func (proxy ManagerProxy) SetName(id int, name string) bool {
	param := make([]interface{}, 2)
	param[0] = id
	param[1] = name
	rq := utils.Request{Op: "SetName", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(bool)
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
func (proxy ManagerProxy) SetAge(id int, age int) bool {
	param := make([]interface{}, 2)
	param[0] = id
	param[1] = age
	rq := utils.Request{Op: "SetAge", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(bool)
}

// SetGender is a function for setting a person is name.
//
// Parameters:
//  id     - the id of the person.
//  gender - the gender of the person.
//
// Returns:
//  a flag if went ok.
//
func (proxy ManagerProxy) SetGender(id int, gender string) bool {
	param := make([]interface{}, 2)
	param[0] = id
	param[1] = gender
	rq := utils.Request{Op: "SetGender", Params: param}
	inv := utils.Invocation{AOR: proxy.AOR, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	return reply[0].(bool)
}
