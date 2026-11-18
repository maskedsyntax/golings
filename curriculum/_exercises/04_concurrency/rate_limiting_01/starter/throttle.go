// rate_limiting_01: One request per tick
//
// Throttle doesn't limit anything yet. See README.md.
//
// Reference: https://gobyexample.com/rate-limiting
package throttle

import "time"

// Throttle calls handle for each request in order, at most once per every.
func Throttle(reqs []string, every time.Duration, handle func(string)) {
	for _, r := range reqs {
		handle(r)
	}
}
