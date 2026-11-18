// stateful_goroutines_01: One goroutine owns the state
//
// Implement NewStore, Get and Set without a mutex. See README.md.
//
// Reference: https://gobyexample.com/stateful-goroutines
package store

type readOp struct {
	key  string
	resp chan int
}

type writeOp struct {
	key  string
	val  int
	resp chan struct{}
}

// Store is a key/value store whose map is owned by a single goroutine.
type Store struct {
	reads  chan readOp
	writes chan writeOp
}

// NewStore creates the store and starts the goroutine that owns its map.
func NewStore() *Store {
	// TODO
	return &Store{}
}

// Get returns the value for key (0 if unset).
func (s *Store) Get(key string) int {
	// TODO
	return 0
}

// Set stores val under key. When Set returns, the write is visible to Get.
func (s *Store) Set(key string, val int) {
	// TODO
}
