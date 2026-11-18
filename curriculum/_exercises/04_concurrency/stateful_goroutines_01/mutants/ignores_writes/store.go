// Mutant: acknowledges writes without applying them.
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

type Store struct {
	reads  chan readOp
	writes chan writeOp
}

func NewStore() *Store {
	s := &Store{reads: make(chan readOp), writes: make(chan writeOp)}
	go func() {
		state := map[string]int{}
		for {
			select {
			case r := <-s.reads:
				r.resp <- state[r.key]
			case w := <-s.writes:
				w.resp <- struct{}{}
			}
		}
	}()
	return s
}

func (s *Store) Get(key string) int {
	resp := make(chan int)
	s.reads <- readOp{key, resp}
	return <-resp
}

func (s *Store) Set(key string, val int) {
	resp := make(chan struct{})
	s.writes <- writeOp{key, val, resp}
	<-resp
}
