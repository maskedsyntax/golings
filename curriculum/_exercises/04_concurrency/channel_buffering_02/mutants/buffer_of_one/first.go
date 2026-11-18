// Mutant: a buffer of 1 still leaks all but one loser.
package first

func First(sources ...func() string) string {
	ch := make(chan string, 1)
	for _, src := range sources {
		go func() { ch <- src() }()
	}
	return <-ch
}
