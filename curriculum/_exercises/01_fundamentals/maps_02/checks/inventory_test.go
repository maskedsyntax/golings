package inventory

import (
	"reflect"
	"testing"
)

func TestSell(t *testing.T) {
	stock := map[string]int{"apple": 3, "pear": 1}
	if !Sell(stock, "apple", 2) || stock["apple"] != 1 {
		t.Fatalf("after Sell(apple, 2) stock = %v, want apple:1", stock)
	}
	if Sell(stock, "apple", 5) || stock["apple"] != 1 {
		t.Fatalf("Sell(apple, 5) with 1 in stock must fail and leave stock unchanged; stock = %v", stock)
	}
	if !Sell(stock, "pear", 1) {
		t.Fatal("Sell(pear, 1) with 1 in stock should succeed")
	}
	if _, ok := stock["pear"]; ok {
		t.Errorf("pear is sold out but still a key: %v", stock)
	}
	if Sell(stock, "kiwi", 1) {
		t.Error("selling an unknown item should fail")
	}
	if _, ok := stock["kiwi"]; ok {
		t.Errorf("a failed sale added a key: %v", stock)
	}
	if Sell(stock, "apple", 0) || Sell(stock, "apple", -1) {
		t.Error("selling zero or negative units should fail")
	}
}

func TestItemsSorted(t *testing.T) {
	stock := map[string]int{}
	names := []string{"plum", "apple", "fig", "date", "kiwi", "banana", "cherry", "grape", "lime", "mango"}
	for _, n := range names {
		stock[n] = 1
	}
	want := []string{"apple", "banana", "cherry", "date", "fig", "grape", "kiwi", "lime", "mango", "plum"}
	// Map order varies between runs, so check repeatedly.
	for range 20 {
		if got := Items(stock); !reflect.DeepEqual(got, want) {
			t.Fatalf("Items = %v, want %v", got, want)
		}
	}
}

func TestItemsEmpty(t *testing.T) {
	if got := Items(map[string]int{}); got == nil || len(got) != 0 {
		t.Errorf("Items(empty) = %#v, want []string{}", got)
	}
}
