// Mutant: replies but forgets to delete.
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

type Store struct {
	reads   chan readOp
	writes  chan writeOp
	deletes chan deleteOp
}

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
				op.resp <- struct{}{}
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

func (s *Store) Delete(key string) {
	resp := make(chan struct{})
	s.deletes <- deleteOp{key, resp}
	<-resp
}
