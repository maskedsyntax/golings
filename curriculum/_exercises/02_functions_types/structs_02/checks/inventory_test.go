package inventory

import (
	"maps"
	"testing"
)

var stock = []Item{
	{Name: "pen", Price: 120, Qty: 10},
	{Name: "book", Price: 1500, Qty: 1},
	{Name: "clip", Price: 5, Qty: 200},
	{Name: "pin", Price: 5, Qty: 50},
}

func TestTotal(t *testing.T) {
	if got := Total(stock); got != 120*10+1500+5*200+5*50 {
		t.Errorf("Total(stock) = %d, want %d", got, 120*10+1500+5*200+5*50)
	}
	if got := Total(nil); got != 0 {
		t.Errorf("Total(nil) = %d, want 0", got)
	}
}

func TestCheapest(t *testing.T) {
	got, ok := Cheapest(stock)
	if !ok || got != stock[2] {
		t.Errorf("Cheapest(stock) = %+v, %v, want %+v, true", got, ok, stock[2])
	}
	if got, ok := Cheapest([]Item{{Name: "only", Price: 900, Qty: 1}}); !ok || got.Name != "only" {
		t.Errorf("Cheapest of one item = %+v, %v", got, ok)
	}
	if got, ok := Cheapest(nil); ok || got != (Item{}) {
		t.Errorf("Cheapest(nil) = %+v, %v, want {}, false", got, ok)
	}
}

func TestByName(t *testing.T) {
	got := ByName(append(stock, Item{Name: "pen", Price: 99, Qty: 3}))
	want := map[string]Item{
		"pen":  {Name: "pen", Price: 99, Qty: 3},
		"book": stock[1],
		"clip": stock[2],
		"pin":  stock[3],
	}
	if !maps.Equal(got, want) {
		t.Errorf("ByName = %+v, want %+v", got, want)
	}
	if m := ByName(nil); m == nil || len(m) != 0 {
		t.Errorf("ByName(nil) = %#v, want an empty, non-nil map", m)
	}
}
