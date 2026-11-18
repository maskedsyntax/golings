# channels_03: Results in input order

`FetchAll` calls `fetch` for every URL concurrently and collects the bodies through a channel.
It should return `results[i] == fetch(urls[i])`. When the fetches finish at different speeds,
the results come back shuffled.

## Your task

Fix `FetchAll` so the results line up with the input. Keep the fetches concurrent, and keep
the channel (don't switch to writing a shared slice from the goroutines).

Reference: https://gobyexample.com/channels
