package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_getObj(t *testing.T) {

	t.Run("compare", func(t *testing.T) {

		want := Obj{
			ID:   1,
			Name: "name",
			Type: "type",
		}

		if got := getObj(2, "name", "type"); !cmp.Equal(got, want) {
			t.Errorf("getObj() => %v", cmp.Diff(got, want))
		}
	})

}
