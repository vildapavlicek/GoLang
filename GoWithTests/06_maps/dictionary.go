package main

const (
	ErrNotFound         = DictionaryErr("not found")
	ErrWordExists       = DictionaryErr("word already in dictionary")
	ErrWorDoesNotdExist = DictionaryErr("word does not exist in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	value, found := d[term]

	if found {
		return value, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWorDoesNotdExist
	case nil:
		d[key] = value
		return nil
	default:
		return err

	}

}

func (d Dictionary) Delete(key string)  {
	delete(d, key)
}

func main() {

}
