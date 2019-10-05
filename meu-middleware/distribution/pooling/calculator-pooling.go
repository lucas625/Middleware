package pooling

import (
	"log"
)

// Pool is a struct to work with many Servants.
//
// Members:
//  Servants - list of Calculator objects.
//  CurrentSize - total number of used Servants.
//
type Pool struct {
	Servants    []interface{}
	CurrentSize int
}

// AddToPool is a function to add a servant to the pool.
//
// Parameters:
//  serv - an object of the type of the pool servants.
//
// Returns:
//  none
//
func (cPool *Pool) AddToPool(serv interface{}) {
	if len(cPool.Servants) <= cPool.CurrentSize {
		log.Fatalln("Reached pool limit.")
	} else {
		cPool.Servants[cPool.CurrentSize] = serv
		cPool.CurrentSize++
	}
}

// GetFromPool is a function to get a servant from the pool.
//
// Parameters:
//  none
//
// Returns:
//  the servant.
//
func (cPool *Pool) GetFromPool() interface{} {
	if cPool.CurrentSize <= 0 {
		log.Fatalln("Empty pool.")
		return nil
	}
	servHolder := cPool.Servants[0]
	cPool.Servants[0] = cPool.Servants[cPool.CurrentSize-1]
	cPool.CurrentSize--
	return servHolder
}

// Initpool is a function to initialize a pool.
//
// Parameters:
//  size - the pool size.
//
// Returns:
//  the pool.
//
func Initpool(size int) *Pool {
	if size > 0 {
		var servants []interface{}
		servants = make([]interface{}, size)
		calcP := Pool{Servants: servants, CurrentSize: 0}
		return &calcP
	}
	return nil
}
