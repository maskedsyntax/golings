# range_over_channels_03: A pipeline that finishes

`Count(Filter(Generate(...), isEven))` should count the even numbers. Each stage ranges over
the previous one's output. One stage never tells the next that it is finished, so the
pipeline hangs.

## Your task

Find the stage that doesn't close its output and fix it.

Reference: https://gobyexample.com/range-over-channels
