# rate_limiting_01: One request per tick

**Rate limiting** controls how often something may happen, for example to protect a server or
to stay within an API quota. The simplest limiter is a ticker:

```go
limiter := time.Tick(200 * time.Millisecond)
for req := range requests {
	<-limiter // wait for the next tick
	fmt.Println("request", req, time.Now())
}
```

## Your task

`Throttle` handles requests in order, but it handles them all at once. Make it handle **at
most one request per `every`**, waiting for a tick before each one, including the first. Use a
`time.Ticker` and stop it when you're done.

Reference: https://gobyexample.com/rate-limiting
