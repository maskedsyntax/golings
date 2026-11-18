// stateful_goroutines_01: One goroutine owns the state (reference solution)
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
	s := &Store{reads: make(chan readOp), writes: make(chan writeOp)}
	go func() {
		state := map[string]int{} // only this goroutine touches state
		for {
			select {
			case r := <-s.reads:
				r.resp <- state[r.key]
			case w := <-s.writes:
				state[w.key] = w.val
				w.resp <- struct{}{}
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

// Set stores val under key. When Set returns, the write is visible to Get.
func (s *Store) Set(key string, val int) {
	resp := make(chan struct{})
	s.writes <- writeOp{key, val, resp}
	<-resp
}
