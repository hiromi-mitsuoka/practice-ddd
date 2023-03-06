package models

import (
	"fmt"
	"reflect"
)

type FullName struct {
	FirstName string
	LastName  string
}

func NewFullName(firstName, lastName string) (FullName, error) {
	if firstName == "" {
		return FullName{}, fmt.Errorf("firstName is required")
	}

	if lastName == "" {
		return FullName{}, fmt.Errorf("lastName is required")
	}

	n := FullName{
		FirstName: firstName,
		LastName:  lastName,
	}

	return n, nil
}

func (f FullName) ChangeFirstName(firstName string) (FullName, error) {
	changedFullName, err := NewFullName(firstName, f.LastName)
	if err != nil {
		return FullName{}, err
	}

	return changedFullName, nil
}

func (f FullName) ChangeLastName(lastName string) (FullName, error) {
	changedFullName, err := NewFullName(f.FirstName, lastName)
	if err != nil {
		return FullName{}, err
	}

	return changedFullName, nil
}

func (f FullName) Equals(otherFullName FullName) bool {
	return reflect.DeepEqual(f, otherFullName)
}
