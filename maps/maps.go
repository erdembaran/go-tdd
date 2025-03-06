package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
    if definition, exists := d[word]; exists {
        return definition, nil
    }
    // errors.New: creates a new errors object
    return "", errors.New("could not find the word you were looking for")
}