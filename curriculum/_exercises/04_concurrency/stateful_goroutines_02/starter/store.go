// stateful_goroutines_02: The missing reply
//
// Delete hangs. Fix it. See README.md.
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

type deleteOp struct {
	key  string
	resp chan struct{}
}

// Store is a key/value store whose map is owned by a single goroutine.
type Store struct {
	reads   chan readOp
	writes  chan writeOp
	deletes chan deleteOp
}

// NewStore creates the store and starts its owner goroutine.
func NewStore() *Store {
	s := &Store{reads: make(chan readOp), writes: make(chan writeOp), deletes: make(chan deleteOp)}
	go func() {
		state := map[string]int{}
		for {
			select {
			case op := <-s.reads:
				op.resp <- state[op.key]
			case op := <-s.writes:
				state[op.key] = op.val
				op.resp <- struct{}{}
			case op := <-s.deletes:
				delete(state, op.key)
			}
		}
	}()
	return s
}

// Get returns the value for key (0 if unset).
func (s *Store) Get(key string) int {
	resp := make(chan int)
	s.reads <- readOp{key, resp}
	return <-resp
}

// Set stores val under key.
func (s *Store) Set(key string, val int) {
	resp := make(chan struct{})
	s.writes <- writeOp{key, val, resp}
	<-resp
}

// Delete removes key.
func (s *Store) Delete(key string) {
	resp := make(chan struct{})
	s.deletes <- deleteOp{key, resp}
	<-resp
}
