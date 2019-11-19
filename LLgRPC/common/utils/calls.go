package utils

import (
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/absoluteobjectreference"
)

// Invocation is a struct for invocation calls.
//
// Members:
//  AOR     - the absolute object reference.
//  Request - the request.
//
type Invocation struct {
	AOR     absoluteobjectreference.AOR
	Request Request
}

// Request is a struct for the request data.
//
// Members:
//  Op     - the operation.
//  Params - list of params.
//
type Request struct {
	Op     string
	Params []interface{}
}

// Reply is a struct for the reply data.
//
// Members:
//  Result - the result of the operation operation.
//
type Reply struct {
	Result []interface{}
}
