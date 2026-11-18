# mutexes_03: The forgotten unlock

`Account` works fine until someone tries to withdraw more than the balance. After that, every
call hangs.

## Your task

Fix `Withdraw`. Make sure the mutex is released on **every** path out of the function.

Reference: https://gobyexample.com/mutexes
