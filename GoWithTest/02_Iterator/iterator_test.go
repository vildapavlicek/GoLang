package iterator

import "testing"

func TestRepeat(t *testing.T) {



	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

t.Run("repeat b 10 times", func(t *testing.T) {
	got := Repeat("b", 10)
	expected := "bbbbbbbbbb"

	assertCorrectMessage(t, expected, got)
})

}


func assertCorrectMessage(t *testing.T, expected, got string){
	if expected != got {
		t.Errorf("expected '%s'; got '%s'", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}