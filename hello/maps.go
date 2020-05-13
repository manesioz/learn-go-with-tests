package main

import "errors"

//Dictionary custom type
type Dictionary map[string]string

//Search function, simple wrapper over map
func (d *Dictionary) Search(key string) (val string, err error) {
	if _, ok := (*d)[key]; ok {
		return (*d)[key], nil
	}
	return "", errors.New("could not find the word you were looking for")
}

//Add function which only adds if key is not present
func (d *Dictionary) Add(key, val string) (finalDict Dictionary, err error) {
	if _, ok := (*d)[key]; ok {
		return nil, errors.New("Error: Word already exists")
	}
	(*d)[key] = val
	return *d, nil
}
