package dictionary

import "testing"

func TestSearch(t *testing.T)  {
	
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		assertDefinition(t, dictionary, "test","this is just a test")
	})
	
	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		
		assertDefinition(t, dictionary, "test","this is just a test")
	})
	
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
	
	t.Run("update word", func(t *testing.T) {
		word := "test"
		definition := "this is just and updated test"
		dictionary := Dictionary{word: definition}
		
		newDefinition := "this is just and updated test"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		
		assertDefinition(t, dictionary, word, newDefinition)
	})
	
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, got := dictionary.Search("test12")

		assertError(t, got, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got.Error() != want.Error() {
		t.Errorf("got %q want %q", got, want)
	}
}