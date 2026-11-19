# testing_and_benchmarking_01: Your first test

In this module **you** write the tests. Go tests live in `_test.go` files and look like this:

```go
func TestSomething(t *testing.T) {
	if got := Something(); got != want {
		t.Errorf("Something() = %v, want %v", got, want)
	}
}
```

`intmin.go` (provided) implements `IntMin`. Your test is in `intmin_test.go`. It doesn't
compile yet, and it only checks one case.

To grade you, golings runs your tests against the real `IntMin`, where they must pass, and then
against broken versions (one always returns its second argument, another returns the larger
number), where they must **fail**.

## Your task

Fix `TestIntMin` and add the case it is missing.

Reference: https://gobyexample.com/testing-and-benchmarking
