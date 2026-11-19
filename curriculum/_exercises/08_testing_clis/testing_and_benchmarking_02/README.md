# testing_and_benchmarking_02: Table-driven tests

Go tests often loop over a table of cases and run each as a *subtest*:

```go
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		if got := F(tt.in); got != tt.want { ... }
	})
}
```

`reverse.go` (provided) implements `Reverse`, which reverses a string by characters (runes),
not bytes.

## Your task

Finish `TestReverse` in `reverse_test.go`:

- Run every case as a subtest with `t.Run`.
- Add enough cases to catch these broken versions: one that reverses **bytes**, one that
  leaves the **first character** in place, and one that **panics on the empty string**.

Reference: https://gobyexample.com/testing-and-benchmarking
