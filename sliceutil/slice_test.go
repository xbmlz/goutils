package sliceutil

import (
	"fmt"
	"testing"
)

type testStruct struct {
	Name string
	Age  int
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3}
	filtered := Filter(slice, func(i int) bool {
		return i == 3
	})
	if filtered[0] != 3 {
		t.Errorf("Expected 3, got %d", filtered[0])
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	listFiltered := Filter(list, func(i testStruct) bool {
		return i.Age > 30
	})

	if len(listFiltered) != 1 {
		t.Errorf("Expected 1, got %d", len(listFiltered))
	}

	if listFiltered[0].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", listFiltered[0].Name)
	}
}

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	mapped := Map(slice, func(i int) int {
		return i * 2
	})
	if mapped[0] != 2 {
		t.Errorf("Expected 2, got %d", mapped[0])
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	listMapped := Map(list, func(i testStruct) string {
		return fmt.Sprintf("%s %d", i.Name, i.Age)
	})

	if len(listMapped) != 3 {
		t.Errorf("Expected 3, got %d", len(listMapped))
	}

	if listMapped[0] != "Alice 25" {
		t.Errorf("Expected Alice 25, got %s", listMapped[0])
	}
}
