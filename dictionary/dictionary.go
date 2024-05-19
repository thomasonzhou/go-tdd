package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

var ErrWordExists = errors.New("word already in dictionary")

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}
