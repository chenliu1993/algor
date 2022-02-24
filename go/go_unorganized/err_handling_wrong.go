package main

import (
	"log"
	"reflect"
)

type E struct{}

func (e *E) Error() string {
	return "initilization"
}

// New news a newly created E instance
func New() E {
	return E{}
}

// GetEmptyErr gets an empty err
func GetEmptyErr() error {
	var e *E = nil
	return e
}

// CheckIfNil checks if the returned E instance is dynamically nil.
func CheckIfNil() string {
	returnMsg := "nil"
	if err := GetEmptyErr(); err != nil {
		log.Println(reflect.TypeOf(err))
		log.Println(reflect.ValueOf(err))
		returnMsg = "not nil"
	}
	return returnMsg
}

func main() {
	log.Println(CheckIfNil())
}
