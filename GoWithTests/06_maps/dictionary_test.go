package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("basic search test", func(t *testing.T) {
		dictionary := Dictionary{
			"test":  "this is just a test",
			"test2": "just another test"}
		/*
			dictionary := map[string]string {
				"test":"this is just a test",
				"test2":"just another test",
			}*/

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("not in dictionary", func(t *testing.T) {
		dictionary := Dictionary{
			"test":  "this is just a test",
			"test2": "just another test",
		}

		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("Expected error")
		}
		assertError(t, err, ErrNotFound)

	})

}

func TestAdd(t *testing.T) {
	t.Run("add entry", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "apple"
		value := "apple is juicy"
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing entry", func(t *testing.T) {
		word := "word"
		definition := "definition"
		dictionary := Dictionary{
			word: definition,
		}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {

		word := "test"
		definition := "sample text"
		newDefinition := "updated sample"
		dictionary := Dictionary{
			word: definition,
		}

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update non-existent word", func(t *testing.T) {
		word := "test"
		definition := "sample definition"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWorDoesNotdExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete entry", func(t *testing.T) {
		word := "test"
		definition := "sample text"
		dictionary := Dictionary{
			word: definition,
		}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("Expected '%s' to be deleted", word)
		}
	})

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %s; want: %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Expected error '%s'; got '%s'", want, got)
	}

}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if got != value {
		t.Errorf("Got '%s'; want '%s'", got, value)
	}
}
