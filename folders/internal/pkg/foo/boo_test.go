package foo

import "testing"

func TestBoo(t *testing.T) {

	t.Run("simple test", func(t *testing.T) {

		want := "foo boo"

		if got := Boo(); got != want {
			t.Errorf("Boo() = %v, want %v", got, want)
		}

	})

}
