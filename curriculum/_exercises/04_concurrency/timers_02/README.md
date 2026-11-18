# timers_02: Debouncing

A **debouncer** waits for a burst of events to settle. It runs its action once, `d` after the
**last** event of the burst. Search boxes use one to avoid querying on every keystroke.

```text
triggers:  x  x  x  x                    x
f runs:                 ·──d──▶ f              ·──d──▶ f
```

This debouncer runs `f` once for **every** trigger.

## Your task

Fix `Debouncer` so `f` runs once, `d` after the last trigger in a burst. A `*time.Timer`'s
`Reset` method reschedules a timer that hasn't fired yet.

Reference: https://gobyexample.com/timers
