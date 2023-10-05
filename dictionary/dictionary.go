package dictionary

import "errors"

type Dictionary map[string]string


var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrAlreadyExists = errors.New("specified key already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error)  {
	val, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}
	
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	
	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = newValue
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string)  {
	delete(d, key)
}