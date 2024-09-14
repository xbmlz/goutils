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

func TestGroupBy(t *testing.T) {
	slice := []int{1, 2, 3, 1, 2, 3}
	grouped := GroupBy(slice, func(i int) int {
		return i
	})
	if len(grouped) != 3 {
		t.Errorf("Expected 3, got %d", len(grouped))
	}
	if len(grouped[1]) != 2 {
		t.Errorf("Expected 2, got %d", len(grouped[1]))
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	listGrouped := GroupBy(list, func(i testStruct) string {
		return i.Name
	})

	if len(listGrouped) != 3 {
		t.Errorf("Expected 3, got %d", len(listGrouped))
	}
	if len(listGrouped["Alice"]) != 2 {
		t.Errorf("Expected 2, got %d", len(listGrouped["Alice"]))
	}
}

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3}
	if Contains(slice, 2) != true {
		t.Errorf("Expected true, got false")
	}
	if Contains(slice, 4) != false {
		t.Errorf("Expected false, got true")
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	if Contains(list, testStruct{"Alice", 25}) != true {
		t.Errorf("Expected true, got false")
	}
	if Contains(list, testStruct{"Dave", 40}) != false {
		t.Errorf("Expected false, got true")
	}
}

func TestIndexOf(t *testing.T) {
	slice := []int{1, 2, 3}
	if IndexOf(slice, 2) != 1 {
		t.Errorf("Expected 1, got %d", IndexOf(slice, 2))
	}
	if IndexOf(slice, 4) != -1 {
		t.Errorf("Expected -1, got %d", IndexOf(slice, 4))
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	if IndexOf(list, testStruct{"Alice", 25}) != 0 {
		t.Errorf("Expected 0, got %d", IndexOf(list, testStruct{"Alice", 25}))
	}
	if IndexOf(list, testStruct{"Dave", 40}) != -1 {
		t.Errorf("Expected -1, got %d", IndexOf(list, testStruct{"Dave", 40}))
	}
}

func TestRemove(t *testing.T) {
	slice := []int{1, 2, 3}
	removed := Remove(slice, 2)
	if len(removed) != 2 {
		t.Errorf("Expected 2, got %d", len(removed))
	}
	if removed[0] != 1 {
		t.Errorf("Expected 1, got %d", removed[0])
	}
	if removed[1] != 3 {
		t.Errorf("Expected 3, got %d", removed[1])
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	removed2 := Remove(list, testStruct{"Alice", 25})
	if len(removed2) != 2 {
		t.Errorf("Expected 2, got %d", len(removed2))
	}
	if removed2[0].Name != "Bob" {
		t.Errorf("Expected Bob, got %s", removed2[0].Name)
	}
	if removed2[1].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", removed2[1].Name)
	}
}

func TestRemoveAll(t *testing.T) {
	slice := []int{1, 2, 3, 1, 2, 3}
	removed := RemoveAll(slice, 2)
	if len(removed) != 4 {
		t.Errorf("Expected 4, got %d", len(removed))
	}
	if removed[0] != 1 {
		t.Errorf("Expected 1, got %d", removed[0])
	}
	if removed[1] != 3 {
		t.Errorf("Expected 3, got %d", removed[1])
	}
	if removed[2] != 1 {
		t.Errorf("Expected 1, got %d", removed[2])
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	removed2 := RemoveAll(list, testStruct{"Alice", 25})
	if len(removed2) != 4 {
		t.Errorf("Expected 4, got %d", len(removed2))
	}
	if removed2[0].Name != "Bob" {
		t.Errorf("Expected Bob, got %s", removed2[0].Name)
	}
	if removed2[1].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", removed2[1].Name)
	}
	if removed2[2].Name != "Bob" {
		t.Errorf("Expected Bob, got %s", removed2[2].Name)
	}
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3}
	Reverse(slice)
	if len(slice) != 3 {
		t.Errorf("Expected 3, got %d", len(slice))
	}
	if slice[0] != 3 {
		t.Errorf("Expected 3, got %d", slice[0])
	}
	if slice[1] != 2 {
		t.Errorf("Expected 2, got %d", slice[1])
	}
	if slice[2] != 1 {
		t.Errorf("Expected 1, got %d", slice[2])
	}

	list := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	Reverse(list)
	if len(list) != 3 {
		t.Errorf("Expected 3, got %d", len(list))
	}
	if list[0].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", list[0].Name)
	}
	if list[1].Name != "Bob" {
		t.Errorf("Expected Bob, got %s", list[1].Name)
	}
	if list[2].Name != "Alice" {
		t.Errorf("Expected Alice, got %s", list[2].Name)
	}
}
